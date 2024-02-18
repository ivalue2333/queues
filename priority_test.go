package queues

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	PriorityQueueTest1(t)
	PriorityQueueTest2(t)
}

func PriorityQueueTest1(t *testing.T) {
	datas := []int{5, 4, 3, 2, 1, 6, 7, 8}
	pq := NewPriorityQueue[int](len(datas), func(iv, jv int) bool {
		return iv > jv
	})
	pq.Push(datas...)
	pq.Push(9)
	i := len(datas) + 1
	for pq.Len() > 0 {
		v := pq.Pop()
		if v != i {
			t.Errorf("v not match")
		}
		i -= 1
	}
}

type testInfo struct {
	Priority int
	Value    string
}

func newTestInfo(p int, v string) *testInfo {
	return &testInfo{
		Priority: p,
		Value:    v,
	}
}

func PriorityQueueTest2(t *testing.T) {
	items := map[int]string{
		3: "banana",
		2: "apple",
		4: "pear",
		1: "good",
		5: "orange",
	}
	pq := NewPriorityQueue[*testInfo](len(items), func(iv, jv *testInfo) bool {
		return iv.Priority > jv.Priority
	})
	for priority, value := range items {
		pq.Push(newTestInfo(priority, value))
	}
	i := len(items)
	for pq.Len() > 0 {
		v := pq.Pop().Value
		expectV := items[i]
		if expectV != v {
			t.Errorf("v not match")
		}
		i -= 1
	}
}
