package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Error("New queue is nil")
	}
}

func TestQueue(t *testing.T) {
	q := NewQueue()

	length := 4
	slc := make([]interface{}, length)
	slc[0] = "string"
	slc[1] = 10
	slc[2] = 0.1
	slc[3] = struct{}{}

	for i := 0; i < length; i++ {
		q.Push(NewTask(slc[i]))
	}

	for i := 0; i < length; i++ {
		tsk := q.Pop()
		if tsk.Value() != slc[i] {
			t.Error("Value does not match")
		}
	}
}

func TestQueue_Length(t *testing.T) {
	q := NewQueue()

	if q.Length() != 0 {
		t.Error("Length does not match")
	}

	q.Push(NewTask(1))

	if q.Length() != 1 {
		t.Error("Length does not match")
	}

	q.Push(NewTask(2))
	if q.Length() != 2 {
		t.Error("Length does not match")
	}

	q.Pop()
	if q.Length() != 1 {
		t.Error("Length does not match")
	}

	q.Push(NewTask(3))
	if q.Length() != 2 {
		t.Error("Length does not match")
	}

	q.Push(NewTask(4))
	q.Push(NewTask(5))
	q.Push(NewTask(6))
	if q.Length() != 5 {
		t.Error("Length does not match")
	}

	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	if q.Length() != 0 {
		t.Error("Length does not match")
	}

	q.Pop()
	if q.Length() != 0 {
		t.Error("Length does not match")
	}

}
