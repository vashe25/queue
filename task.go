package queue

type Task struct {
	element interface{}
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

func (t *Task) Link(pointer *Task) {
	t.next = pointer
}

func (t *Task) Unlink() {
	t.next = nil
}
