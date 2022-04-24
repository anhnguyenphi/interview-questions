package selling_stock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	maxCur, maxSorFar := 0, 0
	for i := 1; i < len(prices); i++ {
		maxCur = max(0, maxCur+(prices[i]-prices[i-1]))
		maxSorFar = max(maxCur, maxSorFar)
	}
	return maxSorFar
}

func TestMaxProfit1(t *testing.T) {
	assert.Equal(t, maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
}

func TestMaxProfit2(t *testing.T) {
	assert.Equal(t, maxProfit([]int{7, 6, 4, 3, 1}), 0)
}
