package a0502_ipo

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	projects := make([][2]int, len(Profits))
	for i := 0; i < len(projects); i++ {
		projects[i] = [2]int{Profits[i], Capital[i]}
	}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i][1] < projects[j][1]
	})
	var ih = &IntHeap{}
	heap.Init(ih)
	j := 0
	for i := 0; i < k; i++ {
		for j < len(Profits) && projects[j][1] <= W {
			heap.Push(ih, projects[j][0])
			j++
		}

		if ih.Len() > 0 {
			W += (heap.Pop(ih)).(int)
		}
	}
	return W
}

type IntHeap []int

func (ih IntHeap) Len() int {
	return len(ih)
}

func (ih IntHeap) Less(i, j int) bool {
	return ih[i] > ih[j]
}

func (ih IntHeap) Swap(i, j int) {
	ih[i], ih[j] = ih[j], ih[i]
}

func (ih *IntHeap) Push(x interface{}) {
	*ih = append(*ih, x.(int))
}

func (ih *IntHeap) Pop() interface{} {
	pre := *ih
	n := len(pre)
	x := pre[n-1]
	*ih = pre[0 : n-1]
	return x
}
