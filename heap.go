package main

import (
	"container/heap"
	"time"
)

type element struct {
	TaskId    string
	Expire    time.Time
	// ...
}

type eleHeap []*element

func (h eleHeap) Len() int           { return len(h) }
func (h eleHeap) Less(i, j int) bool { return h[i].Expire.Before(h[j].Expire) }
func (h eleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *eleHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*element))
}

func (h *eleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]		// 为什么Pop是取slice最末元素？
	*h = old[0 : n-1]   // 难道Golang的堆不是最小堆吗？
	return x
}

func main() {
	to := time.NewTimer(WaitInterval)
	hp := &eleHeap{} // 定义变量
	heap.Init(hp)	 // 堆结构初始化
	for {
		select {
		case ele := <-TaskChan:
			heap.Push(hp, ele) // 入堆
			to.Reset(0)
		case <-to.C:
			for hp.Len() != 0 {
				ele, ok := heap.Pop(hp).(*element) // 出堆
				if ok {
					if time.Now().Before(ele.Expire) {
						heap.Push(hp, ele) // 时辰未到，再次入堆
						to.Reset(ele.Expire.Sub(now))
						break
					}
					// time expired, do task
					// ...
				}
			}
		}
	}
}
