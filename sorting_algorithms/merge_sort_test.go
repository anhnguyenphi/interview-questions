package sorting_algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func merge(arr *[]int, left, middle, right int) {
	n1, n2 := middle-left+1, right-middle
	tmp1, tmp2 := make([]int, n1), make([]int, n2)

	// copy values from arr to tmp1, tmp2
	for i := 0; i < n1; i++ {
		tmp1[i] = (*arr)[left+i]
	}
	for i := 0; i < n2; i++ {
		tmp2[i] = (*arr)[middle+1+i]
	}

	// merge
	var i, j, k = 0, 0, left

	for i < n1 && j < n2 {
		if tmp1[i] < tmp2[j] {
			(*arr)[k] = tmp1[i]
			i++
		} else {
			(*arr)[k] = tmp2[j]
			j++
		}
		k++
	}

	// for the rest of tmp1, tmp2
	for i < n1 {
		(*arr)[k] = tmp1[i]
		i++
		k++
	}

	for j < n2 {
		(*arr)[k] = tmp2[j]
		j++
		k++
	}
}

func mergeSort(arr *[]int, left, right int) []int {
	if left < right {
		middle := int((left + right) / 2)
		mergeSort(arr, left, middle)
		mergeSort(arr, middle+1, right)

		merge(arr, left, middle, right)
	}

	return *arr
}

func TestMerge(t *testing.T) {
	assert.Equal(t, mergeSort(&[]int{3, 1, 2, 4, 3}, 0, 4), []int{1, 2, 3, 3, 4})
}
