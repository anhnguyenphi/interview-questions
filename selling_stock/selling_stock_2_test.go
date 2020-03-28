package selling_stock

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func maxProfit2(prices []int) int {
	maxCur, maxSorFar := 0, 0
	for i := 1; i < len(prices); i++ {
		maxCur = max(0, (prices[i] - prices[i-1]))
		maxSorFar += maxCur
	}
	return maxSorFar
}

func TestMaxProfit2_1(t *testing.T) {
	assert.Equal(t, maxProfit2([]int{7, 1, 5, 3, 6, 4}), 7)
}

func TestMaxProfit2_2(t *testing.T) {
	assert.Equal(t, maxProfit2([]int{1, 2, 3, 4, 5}), 4)
}

func TestMaxProfit2_3(t *testing.T) {
	assert.Equal(t, maxProfit2([]int{7, 6, 4, 3, 1}), 0)
}
