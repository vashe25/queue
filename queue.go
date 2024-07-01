package queue

import "sync"

type Queue struct {
	length int
	first  *Task
	last   *Task
	mtx    sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		length: 0,
		mtx:    sync.Mutex{},
	}
}

func (q *Queue) Push(item *Task) {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	if q.first == nil && q.last == nil {
		q.first = item
		q.last = item
		q.length++
		return
	}

	q.last.Link(item)
	q.last = item

	q.length++
}

func (q *Queue) Pop() *Task {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	if q.first == nil {
		return nil
	}
	if q.first == q.last {
		q.last = nil
	}

	next := q.first.next
	item := q.first
	item.Unlink()
	q.first = next
	q.length--
	return item
}
