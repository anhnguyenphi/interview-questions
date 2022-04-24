package string_problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.geeksforgeeks.org/length-of-the-longest-substring-without-repeating-characters/

func longestUniqueSubsttr(s string) int {
	maxLen := 0
	curLen := 0
	visited := map[byte]int{}

	for i := 0; i < len(s); i++ {
		lastIdx, ok := visited[s[i]]

		if !ok || i-curLen > lastIdx {
			curLen++
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}
			curLen = i - lastIdx
		}

		visited[s[i]] = i
	}

	return maxLen
}

func TestLongestUniqueSubsttr1(t *testing.T) {
	assert.Equal(t, longestUniqueSubsttr("ABDEFGABEF"), 6)
}

func TestLongestUniqueSubsttr2(t *testing.T) {
	assert.Equal(t, longestUniqueSubsttr("GEEKSFORGEEKS"), 7)
}
