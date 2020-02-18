package maximal_rectangle

import (
	"fmt"
	"math"
	"testing"

	"github.com/magiconair/properties/assert"
)

// link: https://leetcode.com/problems/maximal-rectangle/
// link: https://www.geeksforgeeks.	org/given-matrix-o-x-find-largest-subsquare-surrounded-x/

func buildHor(matrix [][]string) [][]int {
	hor := make([][]int, len(matrix))
	for i, row := range matrix {
		hor[i] = make([]int, len(row))
		sum := 0
		for j, val := range row {
			if val == "1" {
				sum++
				hor[i][j] = sum
			} else {
				sum = 0
			}
		}
	}

	return hor
}

func buildVer(matrix [][]string) [][]int {
	ver := make([][]int, len(matrix))
	for i, row := range matrix {
		ver[i] = make([]int, len(row))
		for j, val := range row {
			if val == "1" {
				if i == 0 {
					ver[i][j] = 1
				} else {
					ver[i][j] = ver[i-1][j] + 1
				}
			} else {
				ver[i][j] = 0
			}
		}
	}

	return ver
}

func maximalSubSquare(matrix [][]string) int {
	hor := buildHor(matrix)
	printResult(hor)
	ver := buildVer(matrix)
	printResult(ver)
	N := len(matrix)
	M := len(matrix[0])
	max := 0
	for i := N - 1; i > 0 ; i-- {
		for j := M - 1; j > 0 ; j-- {
			if hor[i][j] > 0 {
				edgeLength := int(math.Min(float64(hor[i][j]), float64(ver[i][j])))
				upEdgeIdx := i - edgeLength + 1
				if upEdgeIdx > -1 &&
					ver[i][j - edgeLength + 1] >= edgeLength {
					if edgeLength > max {
						max = edgeLength
					}
				}
			}
		}
	}
	return max
}

func printResult(nodes [][]int) {
	for _, row := range nodes {
		fmt.Println(row)
	}
	fmt.Println("----------")
}

func TestCase1(t *testing.T)  {
	assert.Equal(t, maximalSubSquare([][]string{
		{"1","1","1","0","0"},
		{"1","0","1","1","1"},
		{"1","1","1","1","1"},
		{"1","0","0","1","0"},
	}), 3)
}

func TestCase2(t *testing.T)  {
	assert.Equal(t, maximalSubSquare([][]string{
		{"1","1","1","1","1"},
		{"0","1","0","1","1"},
		{"1","1","0","1","1"},
		{"1","1","1","1","1"},
	}), 4)
}


