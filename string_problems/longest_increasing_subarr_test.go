package string_problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.geeksforgeeks.org/longest-increasing-subsequence-dp-3/

func lis(arr []int, n int) int {
	if n == 1 {
		return 1
	}

	store := make([]int, n)
	store[0] = 1

	for i := 1; i < n; i++ {
		store[i] = 1
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && store[j]+1 > store[i] {
				store[i] = store[j] + 1
			}
		}
	}

	max := 1
	for _, val := range store {
		if val > max {
			max = val
		}
	}

	return max
}

func TestLIS_1(t *testing.T) {
	assert.Equal(t, lis([]int{3, 10, 2, 1, 20}, 5), 3)
}
