package queues

import "sync"

func NewSyncPriorityQueue[T any](cap int, lessFunc func(iv, jv T) bool) *SyncPriorityQueue[T] {
	return &SyncPriorityQueue[T]{
		q: NewPriorityQueue[T](cap, lessFunc),
	}
}

type SyncPriorityQueue[T any] struct {
	mutex sync.Mutex
	q     *PriorityQueue[T]
}

func (pq *SyncPriorityQueue[T]) Push(v ...T) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()
	pq.q.Push(v...)
}

func (pq *SyncPriorityQueue[T]) Pop() (zero T, ok bool) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()
	if pq.q.Len() == 0 {
		return
	}
	return pq.q.Pop(), true
}

func (pq *SyncPriorityQueue[T]) Len() int {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()
	return pq.q.Len()
}
