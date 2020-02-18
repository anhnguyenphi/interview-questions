package minimum_path_sum

import (
	"math"
	"testing"

	"github.com/magiconair/properties/assert"
)
// link : https://leetcode.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	for i := 0; i < n; i++ {
		for j, val := range grid[i] {
			if i == 0 {
				if j == 0 {
					continue
				} else {
					grid[i][j] = grid[i][j-1] + val
				}
			} else {
				if j == 0 {
					grid[i][j] = grid[i - 1][j] + val
				} else {
					grid[i][j] = int(math.Min(float64(grid[i - 1][j] + val), float64(grid[i][j - 1] + val)))
				}
			}
		}
	}

	return grid[n - 1][m - 1]
}

func TestCase1(t *testing.T) {
	assert.Equal(t, minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}), 7)
}
