package sorting_algorithms

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// link: https://www.geeksforgeeks.org/quick-sort/
// complexity:
// + worse case: O(n^2)
// + best case: O(NlogN)
// pivot is a k largest number after partition process finish

func swap(a, b int) (int, int) {
	return b, a
}

func quickSorr(arr *[]int, low, high int) []int {
	if low < high {
		pivot := low
		for i := low; i < high; i++ {
			if (*arr)[i] < (*arr)[high] {
				(*arr)[i], (*arr)[pivot] = swap((*arr)[i], (*arr)[pivot])
				pivot++
			}
		}

		(*arr)[pivot], (*arr)[high] = swap((*arr)[pivot], (*arr)[high])

		quickSorr(arr, low, pivot-1)
		quickSorr(arr, pivot+1, high)
	}

	return *arr
}

func TestQuickSort(t *testing.T) {
	assert.Equal(t, quickSorr(&[]int{3, 1, 2, 4, 3}, 0, 4), []int{1, 2, 3, 3, 4})
}
