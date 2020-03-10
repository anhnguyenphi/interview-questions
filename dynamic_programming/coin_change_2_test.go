package dynamic_programming

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// https://www.geeksforgeeks.org/coin-change-dp-7/

func coinChange(v int, s []int) int {
	if v == 0 {
		return 1
	}

	result := 0
	for i := 0; i < len(s); i++ {
		if v-s[i] >= 0 {
			result += coinChange(v-s[i], s[i:])
		}
	}

	return result
}

func coinChange2(v int, s []int) int {
	dp := make([][]int, v+1)
	for i := 0; i < v+1; i++ {
		dp[i] = make([]int, len(s))
	}

	for i := 0; i < len(s); i++ {
		dp[0][i] = 1
	}

	for i := 1; i < v+1; i++ {
		for j := 0; j < len(s); j++ {
			// include s[j]
			if i-s[j] >= 0 {
				dp[i][j] = dp[i-s[j]][j]
			}
			// exclude s[j]
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[v][len(s)-1]
}

func TestCoinChange_1(t *testing.T) {
	assert.Equal(t, coinChange(4, []int{1, 2, 3}), 4)
}

func TestCoinChange_2(t *testing.T) {
	assert.Equal(t, coinChange(10, []int{2, 5, 3, 6}), 5)
}

func TestCoinChange2_1(t *testing.T) {
	assert.Equal(t, coinChange2(4, []int{1, 2, 3}), 4)
}

func TestCoinChange2_2(t *testing.T) {
	assert.Equal(t, coinChange2(10, []int{2, 5, 3, 6}), 5)
}
