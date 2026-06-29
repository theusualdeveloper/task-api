package handler

import "github.com/theusualdeveloper/task-api/store"

type MockTaskStore struct {
	tasks []store.Task
}

func NewMockTaskStore() *MockTaskStore {
	return &MockTaskStore{
		tasks: []store.Task{},
	}
}

func (mts *MockTaskStore) Add(title string) store.Task {
	task := store.NewTask(1, title)
	mts.tasks = append(mts.tasks, task)
	return task
}
func (mts *MockTaskStore) GetByID(id int) (store.Task, bool) {
	var task store.Task
	var ok bool
	for _, t := range mts.tasks {
		if t.ID == id {
			task = t
			ok = true
			break
		}
	}
	return task, ok
}
func (mts *MockTaskStore) GetAll() []store.Task {
	return mts.tasks
}
func (mts *MockTaskStore) Delete(id int) bool {
	filtered := []store.Task{}
	var deleted bool
	for _, t := range mts.tasks {
		if t.ID == id {
			deleted = true
			continue
		}
		filtered = append(filtered, t)
	}
	mts.tasks = filtered
	return deleted
}
func (mts *MockTaskStore) Update(id int) (store.Task, bool) {
	var ti int
	var found bool
	for i, t := range mts.tasks {
		if t.ID == id {
			ti = i
			found = true
			break
		}
	}
	if !found {
		return store.Task{}, found
	}
	mts.tasks[ti].Done = true
	return mts.tasks[ti], found
}
