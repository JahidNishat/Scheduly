package repository

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID       string `gorm:"type:uuid;primaryKey"`
	Method   string `gorm:"not null"`
	URL      string `gorm:"not null"`
	Body     string
	RunAt    time.Time `gorm:"not null"`
	Status   string    `gorm:"not null"`
	Attempts int       `gorm:"default:0"`
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (t *TaskRepository) Save(task Task) error {
	return t.db.Create(&task).Error
}

func (t *TaskRepository) Get(id string) (*Task, error) {
	var task Task
	err := t.db.Where("id = ?", id).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *TaskRepository) Tasks() map[string]Task {
	tasks := []Task{}
	now := time.Now()
	t.db.Where("run_at <= ? AND status = ?", now, "pending").Find(&tasks)

	result := map[string]Task{}
	for _, t := range tasks {
		result[t.ID] = t
	}
	return result
}

func (t *TaskRepository) UpdateStatus(id, status string, attempts int) {
	t.db.Model(&Task{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":   status,
			"attempts": attempts,
		})
}
