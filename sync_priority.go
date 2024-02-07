package queues

import "sync"

func NewSyncPriorityQueue[T any](cap int) *SyncPriorityQueue[T] {
	return &SyncPriorityQueue[T]{
		q: NewPriorityQueue[T](cap),
	}
}

type SyncPriorityQueue[T any] struct {
	mutex sync.Mutex
	q     *PriorityQueue[T]
}

func (pq *SyncPriorityQueue[T]) Push(p int, v T) {
	pq.PushItem(NewItem(p, v))
}

func (pq *SyncPriorityQueue[T]) PushItem(item *Item[T]) {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()
	pq.q.PushItem(item)
}

func (pq *SyncPriorityQueue[T]) Pop() T {
	return pq.PopItem().value
}

func (pq *SyncPriorityQueue[T]) PopItem() *Item[T] {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()
	return pq.q.PopItem()
}

func (pq *SyncPriorityQueue[T]) Len() int {
	pq.mutex.Lock()
	defer pq.mutex.Unlock()
	return pq.q.Len()
}
