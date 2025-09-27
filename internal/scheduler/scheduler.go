package scheduler

import (
	"fmt"
	"time"

	"github.com/JahidNishat/scheduly/internal/repository"
)

type Scheduler struct {
	repo *repository.TaskRepository
}

func NewScheduler(repo *repository.TaskRepository) *Scheduler {
	return &Scheduler{repo}
}

func (s *Scheduler) Start() {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			s.runPendingTasks()
		}
	}()
}

func (s *Scheduler) runPendingTasks() {
	for id, task := range s.repo.Tasks() {
		runAt, err := time.Parse(time.RFC3339, task.RunAt)
		if err != nil {
			fmt.Println("invalid time for task: ", id)
			continue
		}

		if time.Now().After(runAt) {
			go s.executeTask(task)
		}
	}
}

func (s *Scheduler) executeTask(task repository.Task) {
	fmt.Printf("Executing task %s: %s %s\n ", task.ID, task.Method, task.URL)
}
