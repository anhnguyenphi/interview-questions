package dynamic_programming

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode.com/problems/jump-game/

func canJump1(nums []int) bool {
	reach := make([]bool, len(nums))
	reach[0] = true
	for i := 0; i < len(nums)-1; i++ {
		if !reach[i] {
			continue
		}
		if nums[i]+i >= len(nums)-1 {
			return true
		}
		for j := 0; j <= nums[i]; j++ {
			reach[i+j] = true
		}
	}
	return reach[len(nums)-1]
}

func canJump2(nums []int) bool {
	i := 0
	reach := 0
	for ; i < len(nums) && i <= reach; i++ {
		reach = maxInt(i+nums[i], reach)
	}
	return i == len(nums)
}

func TestCanJump1_1(t *testing.T) {
	assert.Equal(t, canJump1([]int{2, 3, 1, 1, 4}), true)
}

func TestCanJump1_2(t *testing.T) {
	assert.Equal(t, canJump1([]int{3, 2, 1, 0, 4}), false)
}

func TestCanJump2_1(t *testing.T) {
	assert.Equal(t, canJump2([]int{2, 3, 1, 1, 4}), true)
}

func TestCanJump2_2(t *testing.T) {
	assert.Equal(t, canJump2([]int{3, 2, 1, 0, 4}), false)
}
