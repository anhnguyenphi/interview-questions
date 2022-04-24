package _d_grid_updating

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type node struct {
	x int
	y int
}

func (n node) Neighbours(row, col int) []node {
	neighbours := []node{
		{x: n.x - 1, y: n.y},
		{x: n.x + 1, y: n.y},
		{x: n.x, y: n.y + 1},
		{x: n.x, y: n.y - 1},
	}

	validResult := make([]node, 0)
	for _, n := range neighbours {
		if n.x > -1 && n.x < row &&
			n.y > -1 && n.y < col {
			validResult = append(validResult, n)
		}
	}

	return validResult
}

func minimumDays(row, col int, nodes [][]int) int {
	updatedNodes := make([]node, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if nodes[i][j] == 1 {
				updatedNodes = append(updatedNodes, node{x: i, y: j})
			}
		}
	}

	step := 0
	for len(updatedNodes) > 0 {
		newUpdatedNodes := make([]node, 0)
		for _, n := range updatedNodes {
			neighbours := n.Neighbours(row, col)
			for _, ne := range neighbours {
				if nodes[ne.x][ne.y] == 0 {
					nodes[ne.x][ne.y] = 1
					newUpdatedNodes = append(newUpdatedNodes, node{x: ne.x, y: ne.y})
				}
			}
		}
		printResult(nodes)
		updatedNodes = newUpdatedNodes
		step++
	}
	return step - 1
}

func printResult(nodes [][]int) {
	for _, row := range nodes {
		fmt.Println(row)
	}
	fmt.Println("----------")
}

func TestCase1(t *testing.T) {
	assert.Equal(t, 4, minimumDays(5, 5, [][]int{
		{1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}))
}

func TestCase2(t *testing.T) {
	assert.Equal(t, 3, minimumDays(5, 6, [][]int{
		{0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
	}))
}
