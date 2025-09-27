package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JahidNishat/scheduly/internal/handler"
	"github.com/JahidNishat/scheduly/internal/repository"
	"github.com/JahidNishat/scheduly/internal/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	repo := repository.NewTaskRepository()
	svc := service.NewTaskService(repo)
	h := handler.NewTaskHandler(svc)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Scheduly API is running!"))
	})

	r.Post("/tasks", h.CreateTask)

	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
