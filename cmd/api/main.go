package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JahidNishat/scheduly/internal/handler"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Scheduly API is running")
	})

	http.HandleFunc("/tasks", handler.CreateTask)

	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
