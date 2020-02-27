package __sum

import (
	"math"
	"sort"
	"testing"

	"github.com/magiconair/properties/assert"
)

// https://leetcode.com/problems/3sum-closest/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := 0
	distance := 9999999
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[l] + nums[i] + nums[r]
			if s-target > 0 {
				r--
			} else {
				l++
			}
			newDis := int(math.Abs(float64(s - target)))
			if newDis < distance {
				distance = newDis
				result = s
			}
		}
	}

	return result
}

func TestCase1(t *testing.T) {
	assert.Equal(t, threeSumClosest([]int{-1, 2, 1, -4}, 1), 2)
}

func TestCase2(t *testing.T) {
	assert.Equal(t, threeSumClosest([]int{0, 0, 0}, 1), 0)
}

func TestCase3(t *testing.T) {
	assert.Equal(t, threeSumClosest([]int{1, 1, 1, 0}, -100), 2)
}
