package model

import "time"

type TaskStatus string

const (
	Todo       TaskStatus = "TODO"
	InProgress TaskStatus = "IN_PROGRESS"
	Done       TaskStatus = "DONE"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type User struct {
	Email    string
	Name     string
	Password string
	ID       int
}
