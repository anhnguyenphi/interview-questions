package dynamic_programming

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// https://leetcode.com/problems/jump-game-ii

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	reach := make([]int, len(nums))
	reach[0] = 0
	min := 99999999
	for i := 0; i < len(nums)-1; i++ {
		if i != 0 && reach[i] == 0 {
			continue
		}
		if nums[i]+i >= len(nums)-1 {
			val := reach[i] + 1
			if val < min {
				min = val
			}
		}
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			if reach[i+j] != 0 && reach[i+j] < reach[i]+1 {
				continue
			}
			reach[i+j] = reach[i] + 1
		}
	}
	return min
}

func TestJump_1(t *testing.T) {
	assert.Equal(t, jump([]int{2, 3, 1, 1, 4}), 2)
}
