package queue

import (
	"container/list"
)

type Queue struct {
	queue  *list.List
	putCh  chan func()
	stopCh chan struct{}
}

func NewQueue() *Queue {
	qu := &Queue{
		queue:  list.New(),
		putCh:  make(chan func()),
		stopCh: make(chan struct{}),
	}
	//run listener
	go func() {
		for {
			select {
			case newItem := <-qu.putCh:
				qu.queue.PushBack(newItem)
			case <-qu.stopCh:
				return
			}

		}
	}()
	return qu
}

func (q *Queue) Put(fn func()) {
	q.putCh <- fn
}

func (q *Queue) Get() func() {
	item := q.queue.Front()
	if item == nil {
		return nil
	}
	q.queue.Remove(item)
	if fn, ok := item.Value.(func()); ok {
		return fn
	}
	return nil
}

func (q *Queue) Stop() {
	close(q.stopCh)
}
