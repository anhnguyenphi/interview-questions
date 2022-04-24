package maximum_sum_rectangle

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// link: https://www.geeksforgeeks.org/maximum-sum-rectangle-in-a-2d-matrix-dp-27/

func maxSubArr(arr []int) int {
	bestSum, currentSum := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		currentSum = int(math.Max(float64(arr[i]), float64(currentSum+arr[i])))
		bestSum = int(math.Max(float64(currentSum), float64(bestSum)))
	}

	return bestSum
}

func run(N, M int, arr [][]int) int {
	maxSum := -math.MaxInt32
	for left := 0; left < M; left++ {
		for right := left + 1; right < M; right++ {
			tmpSum := make([]int, N)
			for r := 0; r < N; r++ {
				for i := left; i <= right; i++ {
					tmpSum[r] += arr[r][i]
				}
			}

			maxSum = int(math.Max(float64(maxSubArr(tmpSum)), float64(maxSum)))
		}
	}
	return maxSum
}

func TestCase1(t *testing.T) {
	assert.Equal(t, run(4, 5, [][]int{
		{1, 2, -1, -4, -20},
		{-8, -3, 4, 2, 1},
		{3, 8, 10, 1, 3},
		{-4, -1, 1, 7, -6},
	}), 29)
}
