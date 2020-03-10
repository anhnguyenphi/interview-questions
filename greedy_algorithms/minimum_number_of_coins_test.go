package greedy_algorithms

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// https://www.geeksforgeeks.org/greedy-algorithm-to-find-minimum-number-of-coins/

var denomination = []int{1, 2, 5, 10, 20, 50, 100, 500, 1000}

func minCoins(V int) []int {
	result := make([]int, 0)

	for i := len(denomination) - 1; i > -1; i-- {
		for V-denomination[i] >= 0 {
			V = V - denomination[i]
			result = append(result, denomination[i])
		}
		if V == 0 {
			break
		}
	}

	return result
}

func TestMinCoins_1(t *testing.T) {
	assert.Equal(t, minCoins(70), []int{50, 20})
}

func TestMinCoins_2(t *testing.T) {
	assert.Equal(t, minCoins(121), []int{100, 20, 1})
}
