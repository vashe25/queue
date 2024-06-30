package queue

type Task struct {
	element interface{}
	prev    *Task
	next    *Task
}

func NewTask(element interface{}) *Task {
	return &Task{
		element: element,
	}
}

func (t *Task) Value() interface{} {
	return t.element
}
