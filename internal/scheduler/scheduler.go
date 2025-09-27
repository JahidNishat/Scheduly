package scheduler

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/JahidNishat/scheduly/internal/repository"
	"github.com/spf13/viper"
)

var (
	maxRetries = viper.GetInt("scheduler.max_retries")
	retryDelay = viper.GetDuration("scheduler.retry_delay")
)

type Scheduler struct {
	repo *repository.TaskRepository
}

func NewScheduler(repo *repository.TaskRepository) *Scheduler {
	return &Scheduler{repo}
}

func (s *Scheduler) Start() {
	go func() {
		interval := viper.GetDuration("scheduler.tick_interval")
		ticker := time.NewTicker(interval)
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
			if task.Status == "pending" {
				s.repo.UpdateStatus(task.ID, "executing", task.Attempts)
				go s.executeTask(task)
			}
		}
	}
}

func (s *Scheduler) executeTask(task repository.Task) {
	attempt := 0

	for attempt < maxRetries {
		req, err := http.NewRequest(task.Method, task.URL, bytes.NewBufferString(task.Body))
		if err != nil {
			fmt.Println("invalid request: ", task.ID)
			break
		}

		client := http.Client{Timeout: viper.GetDuration("http.timeout")}
		resp, err := client.Do(req)
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			fmt.Println("task successfully executed: ", task.ID)
			s.repo.UpdateStatus(task.ID, "success", attempt+1)
			return
		}

		attempt++
		fmt.Printf("Task %s failed (attempt %d), retrying...\n", task.ID, attempt)
		time.Sleep(retryDelay)
	}

	s.repo.UpdateStatus(task.ID, "failed", attempt)
}
