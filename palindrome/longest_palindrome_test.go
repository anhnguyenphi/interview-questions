package palindrome

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// link: https://leetcode.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	start, maxLen := 0, 0

	for i := 0; i < l-1; i++ {
		start, maxLen = extendPalindrome(s, i, i, start, maxLen)   // odd
		start, maxLen = extendPalindrome(s, i, i+1, start, maxLen) // even
	}

	return string(s[start : start+maxLen])
}

func extendPalindrome(s string, j, k, start, maxLen int) (int, int) {
	for j >= 0 && k < len(s) && s[j] == s[k] {
		j--
		k++
	}

	if maxLen < k-j-1 {
		return j + 1, k - j - 1
	}

	return start, maxLen
}

func TestLongestPalindrome_Case1(t *testing.T) {
	assert.Equal(t, longestPalindrome("babad"), "bab")
}

func TestLongestPalindrome_Case2(t *testing.T) {
	assert.Equal(t, longestPalindrome("cbbd"), "bb")
}
