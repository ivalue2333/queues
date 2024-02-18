package main

import (
	"fmt"
	"time"

	"github.com/ivalue2333/queues"
)

// demoMaxHeap 大顶堆.
func demoMaxHeap() {
	fmt.Println("demoMaxHeap")
	datas := []int{5, 4, 3, 2, 1, 6, 7, 8}
	pq := queues.NewPriorityQueue[int](len(datas), func(iv, jv int) bool {
		return iv > jv
	})
	pq.Push(datas...)
	for pq.Len() > 0 {
		v := pq.Pop()
		fmt.Println(fmt.Sprintf("v is :%d", v))
	}
}

// demoMinHeap 小顶堆.
func demoMinHeap() {
	fmt.Println("demoMinHeap")
	datas := []int{5, 4, 3, 2, 1, 6, 7, 8}
	pq := queues.NewPriorityQueue[int](len(datas), func(iv, jv int) bool {
		return iv < jv
	})
	pq.Push(datas...)
	for pq.Len() > 0 {
		v := pq.Pop()
		fmt.Println(fmt.Sprintf("v is :%d", v))
	}
}

type exampleInfo1 struct {
	data      string
	createdAt time.Time
	other     int
}

func newExampleInfo1(data string, createdAt time.Time, other int) *exampleInfo1 {
	return &exampleInfo1{
		data:      data,
		createdAt: createdAt,
		other:     other,
	}
}

// demoStruct use struct.
func demoStruct() {
	fmt.Println("demoStruct")
	pq := queues.NewPriorityQueue[*exampleInfo1](0, func(iv, jv *exampleInfo1) bool {
		if iv.data == jv.data {
			return iv.createdAt.After(jv.createdAt)
		}
		return iv.data > jv.data
	})
	pq.Push(newExampleInfo1("1", time.Now(), 1))
	pq.Push(newExampleInfo1("2", time.Now(), 21))
	pq.Push(newExampleInfo1("3", time.Now(), 3))
	pq.Push(newExampleInfo1("2", time.Now(), 22))
	for pq.Len() > 0 {
		v := pq.Pop()
		fmt.Println(fmt.Sprintf("v is :%+v", v))
	}
}

func main() {
	demoMaxHeap()
	demoMinHeap()
	demoStruct()
}
