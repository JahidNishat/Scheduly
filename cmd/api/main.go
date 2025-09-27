package main

import (
	"log"
	"net/http"

	"github.com/JahidNishat/scheduly/internal/config"
	"github.com/JahidNishat/scheduly/internal/handler"
	"github.com/JahidNishat/scheduly/internal/repository"
	"github.com/JahidNishat/scheduly/internal/scheduler"
	"github.com/JahidNishat/scheduly/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()
	log.Println("config loaded successfully")

	r := chi.NewRouter()

	repo := repository.NewTaskRepository()
	svc := service.NewTaskService(repo)
	h := handler.NewTaskHandler(svc)

	sched := scheduler.NewScheduler(repo)
	sched.Start()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Scheduly API is running!"))
	})

	r.Post("/tasks", h.CreateTask)

	port := viper.GetString("server.port")
	log.Println("Listening on port: ", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
