package maximum_path_sum_in_matrix

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"gitlab.myteksi.net/gophers/go/dispatcher/grab-x/commons/utils/math"
)

// link: https://www.geeksforgeeks.org/maximum-path-sum-matrix/
func maximumPathSum(N, M int, matrix [][]int) int {
	for i := 1; i < N; i++ {
		for j := 0; j < M; j++ {
			if j > 0 && j < M - 1 {
				matrix[i][j] += math.MaxInt(matrix[i-1][j], matrix[i-1][j-1], matrix[i-1][j+1])
			} else if j == 0 {
				matrix[i][j] += math.MaxInt(matrix[i-1][j], matrix[i-1][j+1])
			} else {
				matrix[i][j] += math.MaxInt(matrix[i-1][j], matrix[i-1][j-1])
			}
		}
	}

	max := 0
	for _, val := range matrix[N-1] {
		max = math.MaxInt(max, val)
	}
	return max
}

func Test_Case1(t *testing.T)  {
	assert.Equal(t, maximumPathSum(4, 6, [][]int{
		{10, 10 ,2, 0, 20, 4},
		{1, 0, 0, 30, 2, 5},
		{0, 10, 4, 0, 2, 0},
		{1, 0, 2, 20, 0, 4},
	}), 74)
}

func Test_Case2(t *testing.T)  {
	assert.Equal(t, maximumPathSum(3, 3, [][]int{
		{1, 2 ,3},
		{9, 8 ,7},
		{4, 5 ,6},
	}), 17)
}
