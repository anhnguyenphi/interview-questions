package boyer_moore_majority_vote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// link: https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm

func vote(numbers []int) int {
	var m, count int
	for _, val := range numbers {
		if count == 0 {
			count = 1
			m = val
		} else if m == val {
			count++
		} else {
			count--
		}
	}
	return m
}

func TestCase1(t *testing.T) {
	assert.Equal(t, vote([]int{1, 2, 1, 3, 1}), 1)
}

func TestCase2(t *testing.T) {
	assert.Equal(t, vote([]int{2, 2, 1, 3, 2}), 2)
}
