package backtracking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// link: https://leetcode.com/problems/permutations/

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < len(candidates); i++ {
		if target-candidates[i] == 0 {
			result = append(result, []int{candidates[i]})
		}
		if target-candidates[i] < 0 {
			continue
		}
		childResult := combinationSum(candidates[i:], target-candidates[i])
		for j := 0; j < len(childResult); j++ {
			actualResult := append(childResult[j], candidates[i])
			result = append(result, actualResult)
		}
	}

	return result
}

func TestCombinationSum(t *testing.T) {
	result := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(result)

	assert.Equal(t, len(result), 2)
}
