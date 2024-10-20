package main

import "container/heap"

type intHeap [][2]vertex

func (h intHeap) Len() int {
	return len(h)
}

func diff2vertex(a, b vertex) int {
	return b.y - a.y
}

func (h intHeap) Less(i, j int) bool {
	return diff2vertex(h[i][0], h[i][1]) < diff2vertex(h[j][0], h[j][1])
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Pop() any {
	if h.Len() == 0 {
		return nil
	}
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.([2]vertex))
}

func initHeap() *intHeap {
	h := &intHeap{}

	heap.Init(h)

	return h
}
