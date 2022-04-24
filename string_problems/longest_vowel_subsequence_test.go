package string_problems

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.geeksforgeeks.org/length-of-the-longest-substring-without-repeating-characters/

const vowels string = "aeiou"

type subsequence struct {
	Completed bool
	Count     int
}

func subLongestVowelSubsequence(s string, p string, store map[string]subsequence, max int) (int, bool) {
	key := p + strconv.Itoa(len(s))
	if val, ok := store[key]; ok {
		return val.Count, val.Completed
	}

	newMax := max
	isVowels := false
	if len(p) == 1 {
		isVowels = true
	}

	for i := 0; i < len(s); i++ {
		if s[i] == p[0] {
			length, isCompletedSub := subLongestVowelSubsequence(s[i+1:], p, store, max+1)
			if isCompletedSub && length > newMax {
				newMax = length
			}
			key := p + strconv.Itoa(len(s[i+1:]))
			if _, ok := store[key]; !ok {
				store[key] = subsequence{Completed: isCompletedSub, Count: length}
			}

			if !isVowels && isCompletedSub {
				isVowels = isCompletedSub
			}
		}

		if len(p) > 1 && s[i] == p[1] {
			length, isCompletedSub := subLongestVowelSubsequence(s[i+1:], p[1:], store, max+1)
			if isCompletedSub && length > newMax {
				newMax = length
			}
			key := p[1:] + strconv.Itoa(len(s[i+1:]))
			if _, ok := store[key]; !ok {
				store[key] = subsequence{Completed: isCompletedSub, Count: length}
			}
			if !isVowels && isCompletedSub {
				isVowels = isCompletedSub
			}
		}
	}
	return newMax, isVowels
}

func longestVowelSubsequence(s string) int32 {
	count, isCompleted := subLongestVowelSubsequence(s, vowels, map[string]subsequence{}, 0)
	if isCompleted {
		return int32(count)
	}
	return 0
}

func TestCase1(t *testing.T) {
	assert.Equal(t, longestVowelSubsequence("aeiaaioooaauuaeiu"), int32(10))
}
