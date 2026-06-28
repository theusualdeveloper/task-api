package store

import (
	"sync"
)

type TaskStore struct {
	l       sync.RWMutex
	tasks   []Task
	counter int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:   []Task{},
		counter: 1,
	}
}

func (ts *TaskStore) Add(title string) Task {
	ts.l.Lock()
	defer ts.l.Unlock()
	task := NewTask(ts.counter, title)
	ts.tasks = append(ts.tasks, task)
	ts.counter++
	return task
}

func (ts *TaskStore) GetAll() []Task {
	ts.l.RLock()
	defer ts.l.RUnlock()
	return ts.tasks
}

func (ts *TaskStore) GetByID(id int) (Task, bool) {
	ts.l.RLock()
	defer ts.l.RUnlock()
	task := Task{}
	var ok bool
	for _, t := range ts.tasks {
		if t.ID == id {
			task = t
			ok = true
			break
		}
	}
	return task, ok
}

func (ts *TaskStore) Delete(id int) bool {
	ts.l.Lock()
	defer ts.l.Unlock()
	filtered := []Task{}
	var deleted bool
	for _, t := range ts.tasks {
		if t.ID == id {
			deleted = true
			continue
		}
		filtered = append(filtered, t)
	}
	ts.tasks = filtered
	return deleted
}

func (ts *TaskStore) Update(id int) (Task, bool) {
	ts.l.Lock()
	defer ts.l.Unlock()
	var ti int
	var found bool
	for i, t := range ts.tasks {
		if t.ID == id {
			ti = i
			found = true
			break
		}
	}
	if !found {
		return Task{}, found
	}
	ts.tasks[ti].Done = true
	return ts.tasks[ti], found
}
