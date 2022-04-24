package sorting_algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func heapify(arr *[]int, n, i int) {
	largest := i
	left := i*2 + 1
	right := i*2 + 2
	if left < n && (*arr)[largest] < (*arr)[left] {
		largest = left
	}
	if right < n && (*arr)[largest] < (*arr)[right] {
		largest = right
	}

	if largest != i {
		(*arr)[largest], (*arr)[i] = swap((*arr)[largest], (*arr)[i])
		heapify(arr, n, largest)
	}
}

func heapSort(arr []int, n int) []int {
	// build heap
	for i := n/2 - 1; i > -1; i-- {
		heapify(&arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		arr[i], arr[0] = swap(arr[i], arr[0])
		heapify(&arr, i, 0)
	}

	return arr
}

func TestHeapSort(t *testing.T) {
	assert.Equal(t, heapSort([]int{3, 1, 2, 4, 3}, 5), []int{1, 2, 3, 3, 4})
}
