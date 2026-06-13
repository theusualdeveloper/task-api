package store_test

import (
	"testing"

	"github.com/theusualdeveloper/task-api/store"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		task store.Task
	}{
		{
			name: "Test 1",
			task: store.Task{
				ID:    1,
				Title: "Task 1",
			},
		},
		{
			name: "Test 2",
			task: store.Task{
				ID:    2,
				Title: "Task 2",
			},
		},
	}
	ts := store.NewTaskStore()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := ts.Add(tt.task.Title)
			if task.Title != tt.task.Title {
				t.Fatalf("want task title: %s, got: %s", tt.task.Title, task.Title)
			}
			if task.ID != tt.task.ID {
				t.Fatalf("want task id: %d, got: %d", tt.task.ID, task.ID)
			}
		})
	}
}

func TestGetByID(t *testing.T) {
	tests := []struct {
		name  string
		tID   int
		found bool
	}{
		{
			name:  "Test 1",
			tID:   1,
			found: true,
		},
		{
			name:  "Test 2",
			tID:   10,
			found: false,
		},
	}
	ts := store.NewTaskStore()
	ts.Add("Task title")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, found := ts.GetByID(tt.tID)
			if tt.found && !found {
				t.Fatal("task not found")
			}
			if !tt.found && found {
				t.Fatal("task must not found")
			}
		})
	}
}
