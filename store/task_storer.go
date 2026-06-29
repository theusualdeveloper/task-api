package store

type TaskStorer interface {
	Add(title string) Task
	GetAll() []Task
	GetByID(id int) (Task, bool)
	Delete(id int) bool
	Update(id int) (Task, bool)
}
