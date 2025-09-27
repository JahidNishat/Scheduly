package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JahidNishat/scheduly/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service}
}

type CreateTaskRequest struct {
	Method string    `json:"method"`
	Url    string    `json:"url"`
	Body   string    `json:"body,omitempty"`
	RunAt  time.Time `json:"run_at"`
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateTask(req.Method, req.Url, req.Body, req.RunAt)
	if err != nil {
		http.Error(w, "could not create task", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"task_id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
