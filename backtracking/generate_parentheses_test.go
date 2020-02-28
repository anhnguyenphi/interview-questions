package backtracking

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

// https://leetcode.com/problems/generate-parentheses/

func generate(s string, left, right, num int) (result []string) {
	if len(s) == num*2 {
		result = append(result, s)
		return
	}

	if left < num {
		result = append(result, generate(s+"(", left+1, right, num)...)
	}

	if right < left {
		result = append(result, generate(s+")", left, right+1, num)...)
	}

	return
}

func generateParenthesis(n int) []string {
	return generate("", 0, 0, n)
}

func TestGenerateParenthesis_Case1(t *testing.T) {
	result := generateParenthesis(3)
	fmt.Println(result)
	assert.Equal(t, len(result), 5)
}
