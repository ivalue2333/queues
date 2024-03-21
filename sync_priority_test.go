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
	pq := NewSyncPriorityQueue[*testInfo](len(items), func(iv, jv *testInfo) bool {
		return iv.Priority > jv.Priority
	})
	for priority, value := range items {
		pq.Push(newTestInfo(priority, value))
	}
	i := len(items)
	for pq.Len() > 0 {
		v, _ := pq.Pop()
		expectV := items[i]
		if expectV != v.Value {
			t.Errorf("v not match")
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
	pq := NewSyncPriorityQueue[*testInfo](len(items), func(iv, jv *testInfo) bool {
		return iv.Priority > jv.Priority
	})
	wg := sync.WaitGroup{}
	for priority, value := range items {
		p1, v1 := priority, value
		wg.Add(1)
		go func() {
			pq.Push(newTestInfo(p1, v1))
			wg.Done()
		}()
	}
	wg.Wait()
	i := len(items)
	for pq.Len() > 0 {
		v, _ := pq.Pop()
		expectV := items[i]
		if expectV != v.Value {
			t.Errorf("v not match")
		}
		i -= 1
	}
}
