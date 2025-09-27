package repository

import (
	"errors"
)

type Task struct {
	ID       string
	Method   string
	URL      string
	Body     string
	RunAt    string
	Status   string
	Attempts int
}

type TaskRepository struct {
	tasks map[string]Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[string]Task),
	}
}

func (t *TaskRepository) Save(task Task) error {
	if task.ID == "" {
		return errors.New("task ID is empty")
	}
	t.tasks[task.ID] = task
	return nil
}

func (t *TaskRepository) Get(id string) (Task, error) {
	task, ok := t.tasks[id]
	if !ok {
		return Task{}, errors.New("task not found")
	}
	return task, nil
}

func (t *TaskRepository) Tasks() map[string]Task {
	return t.tasks
}

func (t *TaskRepository) UpdateStatus(id, status string, attempts int) {
	task, ok := t.tasks[id]
	if !ok {
		return
	}

	task.Status = status
	task.Attempts = attempts
	t.tasks[id] = task
}
