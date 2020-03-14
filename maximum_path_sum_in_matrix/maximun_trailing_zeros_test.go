package maximum_path_sum_in_matrix

import (
	"math"
	"testing"

	"github.com/magiconair/properties/assert"
)

type node struct {
	Two  int
	Five int
}

func getTwo(val int) int {
	count := 0
	for val%2 == 0 {
		val = val / 2
		count++
	}
	return count
}

func getFive(val int) int {
	count := 0
	for val%5 == 0 {
		val = val / 5
		count++
	}
	return count
}

func _buildRightMatrix(matrix [][]int) [][]node {
	right := make([][]node, len(matrix))
	for i := 0; i < len(matrix); i++ {
		right[i] = make([]node, len(matrix[i]))
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			if j == len(matrix[i])-1 {
				right[i][j] = node{Two: getTwo(matrix[i][j]), Five: getFive(matrix[i][j])}
			} else {
				two := int(math.Max(float64(getTwo(matrix[i][j])), float64(right[i][j+1].Two+getTwo(matrix[i][j]))))
				five := int(math.Max(float64(getFive(matrix[i][j])), float64(right[i][j+1].Five+getFive(matrix[i][j]))))
				right[i][j] = node{Two: two, Five: five}
			}
		}
	}

	return right
}

func _buildLeftMatrix(matrix [][]int) [][]node {
	left := make([][]node, len(matrix))
	for i := 0; i < len(matrix); i++ {
		left[i] = make([]node, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if j == 0 {
				left[i][j] = node{Two: getTwo(matrix[i][j]), Five: getFive(matrix[i][j])}
			} else {
				two := int(math.Max(float64(getTwo(matrix[i][j])), float64(left[i][j-1].Two+getTwo(matrix[i][j]))))
				five := int(math.Max(float64(getFive(matrix[i][j])), float64(left[i][j-1].Five+getFive(matrix[i][j]))))
				left[i][j] = node{Two: two, Five: five}
			}
		}
	}

	return left
}

func _buildDownMatrix(matrix [][]int) [][]node {
	down := make([][]node, len(matrix))
	for i := len(matrix) - 1; i >= 0; i-- {
		down[i] = make([]node, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if i == len(matrix)-1 {
				down[i][j] = node{Two: getTwo(matrix[i][j]), Five: getFive(matrix[i][j])}
			} else {
				two := int(math.Max(float64(getTwo(matrix[i][j])), float64(down[i+1][j].Two+getTwo(matrix[i][j]))))
				five := int(math.Max(float64(getFive(matrix[i][j])), float64(down[i+1][j].Five+getFive(matrix[i][j]))))
				down[i][j] = node{Two: two, Five: five}
			}
		}
	}

	return down
}

func Solution(A [][]int) int {
	left := _buildLeftMatrix(A)
	right := _buildRightMatrix(A)
	down := _buildDownMatrix(A)

	max := 0

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			currentTwo := getTwo(A[i][j])
			currentFive := getFive(A[i][j])
			two := int(math.Max(float64(left[i][j].Two+down[i][j].Two-currentTwo), float64(right[i][j].Two+down[i][j].Two-currentTwo)))
			five := int(math.Max(float64(left[i][j].Five+down[i][j].Five-currentFive), float64(right[i][j].Five+down[i][j].Five-currentFive)))
			ten := int(math.Min(float64(two), float64(five)))
			if ten > max {
				max = ten
			}
		}
	}

	return max
}

func TestCase1(t *testing.T) {
	s := Solution([][]int{
		{1, 10, 1, 10},
		{1, 1, 1, 10},
		{1, 1, 1, 1},
	})
	assert.Equal(t, s, 3)
}
