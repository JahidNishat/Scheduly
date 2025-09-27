package service

import (
	"github.com/JahidNishat/scheduly/internal/repository"
	"github.com/google/uuid"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo}
}

func (s *TaskService) CreateTask(method, url, body, runAt string) (string, error) {
	id := uuid.NewString()
	task := repository.Task{
		ID:     id,
		Method: method,
		URL:    url,
		Body:   body,
		RunAt:  runAt,
	}

	if err := s.repo.Save(task); err != nil {
		return "", err
	}

	return id, nil
}
