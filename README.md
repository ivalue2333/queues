# queues

golang 队列/优先队列

# 优先队列/并发安全优先队列

补充 golang 生态的优先队列，代码实现清晰、简单，没有额外的引用。

1. 基于二叉堆实现
2. 泛型
3. 单元测试完整
4. 支持自定义 compare 函数
5. 并发安全
6. API 简单易用

## example

```shell
go get github.com/ivalue2333/queues
```

```go
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

func main() {
	demoMaxHeap()
}
```


