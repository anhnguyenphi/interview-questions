package largest_sum_contiguous_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// link: https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/

func run(arr []int) int {
	N := len(arr)
	maxSoFar, maxEndingHere := -999, 0
	for i := 0; i < N; i++ {
		maxEndingHere += arr[i]
		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}

		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}
	return maxSoFar
}

func Test_Case1(t *testing.T) {
	assert.Equal(t, run([]int{-2, -3, 4, -1, -2, 1, 5, -3}), 7)
}

func Test_Case2(t *testing.T) {
	assert.Equal(t, run([]int{-2, 10, -18, 4, -1, -2, 1, 5, -3}), 10)
}

func Test_Case3(t *testing.T) {
	assert.Equal(t, run([]int{-2, -3}), -2)
}

func Test_Case4(t *testing.T) {
	assert.Equal(t, run([]int{-2, 10, -18, 4, -1, -2, 1, 5, -7, 11, -3}), 11)
}

func Test_Case5(t *testing.T) {
	assert.Equal(t, run([]int{2, -1, 0, 0, -1, 2, 2, -2, 2}), 4)
}
