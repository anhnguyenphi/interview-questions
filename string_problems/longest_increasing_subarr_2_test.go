package string_problems

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// https://www.geeksforgeeks.org/longest-monotonically-increasing-subsequence-size-n-log-n/

func lis2BinarySearch(arr []int, l, r, key int) int {
	for r-l > 1 {
		m := int((r + l) / 2)
		if arr[m] > key {
			r = m
		} else {
			l = m
		}
	}
	return r
}

func lis2(arr []int, n int) int {
	// last Element -> length
	activeList := make([]int, n)
	activeList[0] = arr[0]
	length := 1
	for i := 1; i < n; i++ {
		// new smallest number
		if arr[i] < activeList[0] {
			activeList[0] = arr[i]
		} else if arr[i] > activeList[length-1] { // clone and extend
			activeList[length] = arr[i]
			length++
		} else { // between
			activeList[lis2BinarySearch(activeList, 0, length-1, arr[i])] = arr[i]
		}
	}

	return length
}

func TestLIS2_1(t *testing.T) {
	assert.Equal(t, lis2([]int{3, 10, 2, 1, 20}, 5), 3)
}

func TestLIS2_2(t *testing.T) {
	assert.Equal(t, lis2([]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}, 16), 6)
}
