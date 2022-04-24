package backtracking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/permutations/

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		newNums := make([]int, len(nums)-1)
		for j := 0; j < len(newNums); j++ {
			if j < i {
				newNums[j] = nums[j]
			} else {
				newNums[j] = nums[j+1]
			}
		}
		gen := permute(newNums)
		for j := 0; j < len(gen); j++ {
			result = append(result, append(gen[j], nums[i]))
		}
	}

	return result
}

func TestPermute_Case1(t *testing.T) {
	result := permute([]int{1, 2, 3})
	fmt.Println(result)
	assert.Equal(t, len(result), 6)
}
