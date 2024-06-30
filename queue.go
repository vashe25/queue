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

	q.last.next = item
	item.prev = q.last
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
	if next != nil {
		next.prev = nil
	}
	item := q.first
	item.prev = nil
	item.next = nil
	q.first = next
	q.length--
	return item
}
