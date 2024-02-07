package main

import (
	"fmt"
	"sync"

	"github.com/ivalue2333/queues"
)

func main() {
	items := map[int]string{
		3: "banana",
		2: "apple",
		4: "pear",
		1: "good",
		6: "6",
		5: "orange",
		7: "7",
	}
	pq := queues.NewSyncPriorityQueue[string](len(items))
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
	for pq.Len() > 0 {
		item := pq.PopItem()
		fmt.Println(fmt.Sprintf("item priority:%d, val:%v", item.GetPriority(), item.GetValue()))
	}
}
