package dynamic_programming

import (
	"math"
	"testing"

	"github.com/magiconair/properties/assert"
)

// https://leetcode.com/problems/coin-change/

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	for i := 0; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				dp[i] = int(math.Min(float64(dp[i]), 1+float64(dp[i-coins[j]])))
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func TestCoinChange1_Case1(t *testing.T) {
	assert.Equal(t, coinChange1([]int{186, 419, 83, 408}, 6249), 20)
}
