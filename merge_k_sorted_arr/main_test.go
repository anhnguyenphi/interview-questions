package merge_k_sorted_arr

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

type minHeap struct {
	arr [][]int
}

func (h *minHeap) Push(arr []int) {
	h.arr = append(h.arr, arr)
	h.bottomUp(len(h.arr) - 1)
}

func (h *minHeap) Pop() []int {
	val := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[0 : len(h.arr)-1]
	h.topDown(0)
	return val
}

func (h *minHeap) bottomUp(i int) {
	var parent int
	if i%2 == 0 {
		parent = i/2 - 1
	} else {
		parent = (i - 1) / 2
	}

	if parent < 0 {
		return
	}

	if h.arr[parent][0] > h.arr[i][0] {
		h.arr[parent], h.arr[i] = h.arr[i], h.arr[parent]
		h.bottomUp(parent)
	}
}

func (h *minHeap) topDown(i int) {
	n := len(h.arr)
	smallest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h.arr[smallest][0] > h.arr[left][0] {
		smallest = left
	}

	if right < n && h.arr[smallest][0] > h.arr[right][0] {
		smallest = right
	}

	if smallest != i {
		h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
		h.topDown(smallest)
	}
}

func (h *minHeap) Empty() bool {
	return len(h.arr) == 0
}

func mergeArr(arrays ...[]int) []int {
	k := len(arrays)
	heap := &minHeap{}
	for i := 0; i < k; i++ {
		heap.Push([]int{arrays[i][0], i, 0})
	}

	resutl := make([]int, 0)

	for !heap.Empty() {
		top := heap.Pop()

		val := top[0]
		i := top[1]
		j := top[2]

		resutl = append(resutl, val)

		if j+1 < len(arrays[i]) {
			heap.Push([]int{arrays[i][j+1], i, j + 1})
		}
	}

	return resutl
}

func TestMergeArr(t *testing.T) {
	assert.Equal(t, mergeArr([]int{2, 6, 12}, []int{1, 9}, []int{23, 34, 90, 2000}),
		[]int{1, 2, 6, 9, 12, 23, 34, 90, 2000})
}
