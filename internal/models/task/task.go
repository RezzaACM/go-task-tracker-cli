package models

import "time"

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Completed  Status = "completed"
	All        Status = "all"
)

type TaskManager struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Status      Status    `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
