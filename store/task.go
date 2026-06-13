package store

import "time"

type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

func NewTask(id int, title string) Task {
	return Task{
		ID:        id,
		Title:     title,
		CreatedAt: time.Now(),
	}
}
