package queues

import (
	"sync"
	"testing"
)

func TestSyncPriorityQueue(t *testing.T) {
	SyncPriorityQueueTest1(t)
	SyncPriorityQueueTestCurrent(t)
}

func SyncPriorityQueueTest1(t *testing.T) {
	items := map[int]string{
		3: "banana",
		2: "apple",
		4: "pear",
		1: "good",
		5: "orange",
	}
	pq := NewPriorityQueue[string](len(items))
	for priority, value := range items {
		pq.Push(priority, value)
	}
	i := len(items)
	for pq.Len() > 0 {
		item := pq.PopItem()
		if i != item.priority {
			t.Errorf("i:%d not match priority:%d", i, item.priority)
		}
		i -= 1
	}
}

func SyncPriorityQueueTestCurrent(t *testing.T) {
	items := map[int]string{
		3: "banana",
		2: "apple",
		4: "pear",
		1: "good",
		6: "6",
		5: "orange",
		7: "7",
	}
	pq := NewSyncPriorityQueue[string](len(items))
	wg := sync.WaitGroup{}
	for priority, value := range items {
		p1, v1 := priority, value
		wg.Add(1)
		go func() {
			pq.Push(p1, v1)
			wg.Done()
		}()
	}
	wg.Wait()
	i := len(items)
	for pq.Len() > 0 {
		item := pq.PopItem()
		if i != item.priority {
			t.Errorf("i:%d not match priority:%d", i, item.priority)
		}
		i -= 1
	}
}
