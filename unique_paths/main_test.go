package unique_paths

import (
	"strconv"
	"testing"

	"github.com/magiconair/properties/assert"
)

// https://leetcode.com/problems/unique-paths/

func step(m, n int, store map[string]int) int {
	key := strconv.Itoa(m) + "," + strconv.Itoa(n)
	val, ok := store[key]
	if ok {
		return val
	}

	if m == 1 && n == 1 {
		return 1
	}

	var result int
	if n == 1 {
		result = step(m-1, n, store)
	} else if m == 1 {
		result = step(m, n-1, store)
	} else {
		result = step(m-1, n, store) + step(m, n-1, store)
	}

	store[key] = result
	return result
}

func run(m, n int) int {
	store := map[string]int{}
	return step(m, n, store)
}

func TestCase1(t *testing.T) {
	assert.Equal(t, run(3, 2), 3)
}

func TestCase2(t *testing.T) {
	assert.Equal(t, run(7, 3), 28)
}
