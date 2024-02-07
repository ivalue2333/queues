package main

import (
	"fmt"

	"github.com/ivalue2333/queues"
)

func main() {
	items := map[int]string{
		3: "banana",
		2: "apple",
		4: "pear",
		1: "good",
		5: "orange",
	}
	pq := queues.NewPriorityQueue[string](len(items))
	for priority, value := range items {
		pq.Push(priority, value)
	}
	for pq.Len() > 0 {
		item := pq.PopItem()
		fmt.Println(fmt.Sprintf("item priority:%d, val:%v", item.GetPriority(), item.GetValue()))
	}
}
