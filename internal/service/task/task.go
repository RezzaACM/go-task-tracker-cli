package service

import (
	"encoding/json"
	"os"
	"time"

	models "github.com/yourusername/task-tracer-cli/internal/models/task"
)

type TaskManager struct {
	Tasks []models.Task `json:"tasks"`
}

func (tm *TaskManager) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, tm); err != nil {
		return err
	}
	return nil
}

func (tm *TaskManager) Save(filename string) error {
	data, err := json.MarshalIndent(tm, "", "")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (tm *TaskManager) Add(task models.Task) {
	id := len(tm.Tasks) + 1
	tm.Tasks = append(tm.Tasks, models.Task{
		ID:          id,
		Title:       task.Title,
		Status:      models.Todo,
		Description: task.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
}

func (tm *TaskManager) Update(id int, task models.Task) {
	for i, t := range tm.Tasks {
		if t.ID == id {
			// Retain previous values if not updated
			if task.Title == "" {
				task.Title = t.Title
			}
			if task.Status == "" {
				task.Status = t.Status
			}
			if task.Description == "" {
				task.Description = t.Description
			}
			tm.Tasks[i] = models.Task{
				ID:          id,
				Title:       task.Title,
				Status:      task.Status,
				Description: task.Description,
				CreatedAt:   t.CreatedAt,
				UpdatedAt:   time.Now(),
			}
			return
		}
	}
}

func (tm *TaskManager) Delete(id int) {
	for i, t := range tm.Tasks {
		if t.ID == id {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			return
		}
	}
}

func (tm *TaskManager) List(status models.Status) []models.Task {
	var tasks []models.Task
	for _, t := range tm.Tasks {
		if t.Status == status {
			tasks = append(tasks, t)
		}
		if status == models.All {
			tasks = append(tasks, t)
		}
	}
	return tasks
}

func (tm *TaskManager) SetStatus(id int, status models.Status) {
	for i, t := range tm.Tasks {
		if t.ID == id {
			tm.Tasks[i].Status = status
			tm.Tasks[i].UpdatedAt = time.Now()
			return
		}
	}
}
