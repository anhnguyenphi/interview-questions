package palindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// link: https://leetcode.com/problems/palindromic-substrings/

func countSubstrings(s string) int {
	l := len(s)
	if l < 2 {
		return 1
	}
	result := 0

	for i := 0; i < l; i++ {
		result += extendAndCount(s, i, i) // odd
		if i+1 < l {
			result += extendAndCount(s, i, i+1) // even
		}
	}

	return result
}

func extendAndCount(s string, j, k int) int {
	count := 0
	for j > -1 && k < len(s) && s[j] == s[k] {
		j--
		k++
		count++
	}

	return count
}

func TestCountSubstrings_Case1(t *testing.T) {
	assert.Equal(t, countSubstrings("abc"), 3)
}

func TestCountSubstrings_Case2(t *testing.T) {
	assert.Equal(t, countSubstrings("aaa"), 6)
}
