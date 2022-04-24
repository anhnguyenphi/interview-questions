package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/jump-game-ii

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func jump(nums []int) int {
	steps := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i] && j+i < len(nums); j++ {
			if steps[i+j] == 0 {
				steps[i+j] = steps[i] + 1
			} else {
				steps[i+j] = min(steps[i]+1, steps[i+j])
			}
		}
	}
	return steps[len(nums)-1]
}

func TestJump_1(t *testing.T) {
	assert.Equal(t, jump([]int{2, 3, 1, 1, 4}), 2)
}
