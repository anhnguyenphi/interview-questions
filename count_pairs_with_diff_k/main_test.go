package count_pairs_with_diff_k

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func countPairsWithDiffK(arr []int, k int) [][]int {
	store := map[int]bool{}
	for _, ele := range arr {
		store[ele] = true
	}

	result := make([][]int, 0)

	for _, ele := range arr {
		if _, ok := store[ele - k]; ok {
			result = append(result, []int{ele - k, ele})
			delete(store, ele - k)
		}

		if _, ok := store[ele + k]; ok {
			result = append(result, []int{ele + k, ele})
			delete(store, ele + k)
		}

		delete(store, ele)
	}

	return result
}

func TestCase1(t *testing.T) {
	assert.Equal(t, countPairsWithDiffK([]int{1, 5, 3, 4, 2}, 3), [][]int{
		{4, 1},
		{2, 5},
	})
}