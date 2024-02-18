package queues

type Item[T any] struct {
	value T
	index int
}

func newItem[T any](v T) *Item[T] {
	return &Item[T]{
		value: v,
	}
}

func NewPriorityQueue[T any](cap int, lessFunc func(iv, jv T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		datas:    make([]*Item[T], 0, cap),
		lessFunc: lessFunc,
	}
}

type PriorityQueue[T any] struct {
	datas    []*Item[T]
	lessFunc func(iv, jv T) bool
}

func (pq *PriorityQueue[T]) Push(vv ...T) {
	for _, v := range vv {
		pq.push(v)
	}
}

func (pq *PriorityQueue[T]) push(v T) {
	item := newItem(v)
	n := len(pq.datas)
	item.index = n
	pq.datas = append(pq.datas, item)
	// 排序
	pq.up(pq.Len() - 1)
}

func (pq *PriorityQueue[T]) Pop() T {
	n := pq.Len() - 1
	pq.swap(0, n)
	// 排序
	pq.down(0, n)
	// 返回
	return pq.doPop().value
}

func (pq *PriorityQueue[T]) doPop() *Item[T] {
	old := pq.datas
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	pq.datas = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.datas)
}

func (pq *PriorityQueue[T]) less(i, j int) bool {
	return pq.lessFunc(pq.datas[i].value, pq.datas[j].value)
}

func (pq *PriorityQueue[T]) swap(i, j int) {
	pq.datas[i], pq.datas[j] = pq.datas[j], pq.datas[i]
	pq.datas[i].index = i
	pq.datas[j].index = j
}

func (pq *PriorityQueue[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		j = i
	}
}

func (pq *PriorityQueue[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && pq.less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !pq.less(j, i) {
			break
		}
		pq.swap(i, j)
		i = j
	}
	return i > i0
}
