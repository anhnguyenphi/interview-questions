package maximum_path_sum_in_matrix

import (
	"fannpa/hello/util"
	"math"
	"testing"

	"github.com/magiconair/properties/assert"
)

// find a maximum path in 2d matrix (has negative number) that we can only move left and right 1 time and down 1 time

func buildRightMatrix(matrix [][]int) [][]int {
	right := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		right[i] = make([]int, len(matrix[i]))
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			if j == len(matrix[i])-1 {
				right[i][j] = matrix[i][j]
			} else {
				right[i][j] = int(math.Max(float64(matrix[i][j]), float64(right[i][j+1])+float64(matrix[i][j])))
			}
		}
	}

	return right
}

func buildLeftMatrix(matrix [][]int) [][]int {
	left := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		left[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				left[i][j] = matrix[i][j]
			} else {
				left[i][j] = int(math.Max(float64(matrix[i][j]), float64(left[i][j-1])+float64(matrix[i][j])))
			}
		}
	}

	return left
}

func buildDownMatrix(matrix [][]int) [][]int {
	down := make([][]int, len(matrix))
	for i := len(matrix) - 1; i >= 0; i-- {
		down[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if i == len(matrix)-1 {
				down[i][j] = matrix[i][j]
			} else {
				down[i][j] = int(math.Max(float64(matrix[i][j]), float64(down[i+1][j])+float64(matrix[i][j])))
			}
		}
	}

	return down
}

func maxSumPath(matrix [][]int) int {
	left := buildLeftMatrix(matrix)
	right := buildRightMatrix(matrix)
	down := buildDownMatrix(matrix)

	util.PrintIntMatrix(left)
	util.PrintIntMatrix(right)
	util.PrintIntMatrix(down)

	max := math.Max(float64(left[0][0]+down[0][0]), float64(down[0][0]+right[0][0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			val := math.Max(float64(left[i][j]+down[i][j]-matrix[i][j]), float64(down[i][j]+right[i][j]-matrix[i][j]))
			if val > max {
				max = val
			}
		}
	}
	return int(max)
}

func TestMaxSumPath_1(t *testing.T) {
	assert.Equal(t, maxSumPath([][]int{
		{1, 1, 3},
		{2, -1, 3},
		{-1, 0, 3},
	}), 11)
}
