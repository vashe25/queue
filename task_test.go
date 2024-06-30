package queue

import (
	"reflect"
	"testing"
)

func TestNewTask(t *testing.T) {
	tsk := NewTask("task")
	if tsk == nil {
		t.Error("New task is nil")
	}

	if reflect.TypeOf(tsk).String() != "*queue.Task" {
		t.Error("New task is not a queue.Task")
	}
}
