package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func ConnectDB() *gorm.DB {
	dbOnce.Do(func() {
		host := viper.GetString("database.host")
		port := viper.GetString("database.port")
		user := viper.GetString("database.user")
		password := viper.GetString("database.password")
		name := viper.GetString("database.dbname")
		sslmode := viper.GetString("database.sslmode")

		var err error
		dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, name, sslmode)
		db, err = gorm.Open(postgres.Open(dns))
		if err != nil {
			log.Fatalf("Error connecting to database, %s", err)
		}
	})

	return db
}
