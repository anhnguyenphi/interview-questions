package dynamic_programming

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.geeksforgeeks.org/0-1-knapsack-problem-dp-10/

// Returns the maximum value that
// can be put in a knapsack of capacity W
func knapSack1(W int, wt, val []int, n int) int {
	if W == 0 || n == 0 {
		return 0
	}

	if W-wt[n-1] < 0 {
		return knapSack1(W, wt, val, n-1)
	}

	return int(math.Max(float64(val[n-1])+float64(knapSack1(W-wt[n-1], wt, val, n-1)), float64(knapSack1(W, wt, val, n-1))))
}

func knapSack2(W int, wt, val []int, n int) int {
	dp := make([][]int, W+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < W+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if i >= wt[j-1] {
				dp[i][j] = int(math.Max(float64(dp[i][j-1]), float64(val[j-1])+float64(dp[i-wt[j-1]][j-1])))
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[W][n]
}

func TestKnapSack1Case1(t *testing.T) {
	assert.Equal(t, knapSack1(50, []int{10, 20, 30}, []int{60, 100, 120}, 3), 220)
}

func TestKnapSack2Case1(t *testing.T) {
	assert.Equal(t, knapSack2(50, []int{10, 20, 30}, []int{60, 100, 120}, 3), 220)
}
