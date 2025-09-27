package repository

import (
	"errors"
)

type Task struct {
	ID     string
	Method string
	URL    string
	Body   string
	RunAt  string
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
