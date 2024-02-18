package main

import (
	"fmt"
	"sync"

	"github.com/ivalue2333/queues"
)

// demoMaxHeapSync 大顶堆.
func demoMaxHeapSync() {
	fmt.Println("demoMaxHeapSync")
	datas := []int{5, 4, 3, 2, 1, 6, 7, 8}
	pq := queues.NewSyncPriorityQueue[int](len(datas), func(iv, jv int) bool {
		return iv > jv
	})
	pq.Push(datas...)
	for pq.Len() > 0 {
		v := pq.Pop()
		fmt.Println(fmt.Sprintf("v is :%d", v))
	}
}

// demoMinHeapSync 小顶堆.
func demoMinHeapSync() {
	fmt.Println("demoMinHeapSync")
	datas := []int{5, 4, 3, 2, 1, 6, 7, 8}
	pq := queues.NewSyncPriorityQueue[int](len(datas), func(iv, jv int) bool {
		return iv < jv
	})
	pq.Push(datas...)
	for pq.Len() > 0 {
		v := pq.Pop()
		fmt.Println(fmt.Sprintf("v is :%d", v))
	}
}

func demoSyncCurrent() {
	fmt.Println("demoSyncCurrent")
	size := 1000
	items := make([]int, 0, size)
	for i := 0; i < size; i++ {
		items = append(items, i)
	}
	pq := queues.NewSyncPriorityQueue[int](0, func(iv, jv int) bool {
		return iv > jv
	})
	wg := sync.WaitGroup{}
	for _, v := range items {
		val := v
		wg.Add(1)
		go func() {
			pq.Push(val)
			wg.Done()
		}()
	}
	wg.Wait()
	i := pq.Len() - 1
	for pq.Len() > 0 {
		v := pq.Pop()
		if v != i {
			fmt.Println("not good", v, i)
			return
		}
		i -= 1
	}
}

func main() {
	demoMaxHeapSync()
	demoMinHeapSync()
	demoSyncCurrent()
}
