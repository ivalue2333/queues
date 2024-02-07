package queues

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	PriorityQueueTest1(t)
}

func PriorityQueueTest1(t *testing.T) {
	items := map[int]string{
		3: "banana",
		2: "apple",
		4: "pear",
		1: "good",
		5: "orange",
	}
	pq := NewPriorityQueue[string](len(items))
	for priority, value := range items {
		pq.Push(NewItem(priority, value))
	}
	i := len(items)
	for pq.Len() > 0 {
		item := pq.Pop()
		if i != item.priority {
			t.Errorf("i:%d not match priority:%d", i, item.priority)
		}
		i -= 1
	}
}
