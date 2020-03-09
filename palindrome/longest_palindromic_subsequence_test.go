package palindrome

import (
	"math"
	"testing"

	"github.com/magiconair/properties/assert"
)

// https://leetcode.com/problems/longest-palindromic-subsequence/discuss/99101/Straight-forward-Java-DP-solution

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}

	for i := len(s) - 1; i > -1; i-- {
		dp[i][i] = 1

		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i][j-1])))
			}
		}
	}

	return dp[0][len(s)-1]
}

func TestLongestPalindromeSubseq1(t *testing.T) {
	assert.Equal(t, longestPalindromeSubseq("bbbab"), 4)
}

func TestLongestPalindromeSubseq2(t *testing.T) {
	assert.Equal(t, longestPalindromeSubseq("cbbd"), 2)
}
