package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxArea(height []int) int {
	max := 0
	x := 0
	y := len(height) - 1
	for x != y {
		var area int
		if height[x] > height[y] {
			area = height[y] * (y - x)
			y--
		} else {
			area = height[x] * (y - x)
			x++
		}

		if area > max {
			max = area
		}
	}
	return max
}

func TestMaxArea(t *testing.T) {
	assert.Equal(t, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
}
