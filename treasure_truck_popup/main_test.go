package treasure_truck_popup

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

type node struct {
	x int
	y int
}

func (n node) Neighbours(row, col int) []node {
	neighbours := []node{
		{ x: n.x - 1, y: n.y },
		{ x: n.x + 1, y: n.y },
		{ x: n.x, y: n.y + 1 },
		{ x: n.x, y: n.y - 1},
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

func TestFindNumberOfTruckPopups1(t *testing.T) {
	assert.Equal(t, 2, findNumberOfTruckPopups(4, 4, [][]int{
		{1,1,0,0},
		{0,0,0,0},
		{0,0,1,1},
		{0,0,0,0},
	}))
}
func TestFindNumberOfTruckPopups2(t *testing.T) {
	assert.Equal(t, 7, findNumberOfTruckPopups(7,7, [][]int{
		{1,0,0,0,0,0,0},
		{0,1,0,0,0,0,0},
		{0,0,1,0,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,0,1,0,0},
		{0,0,0,0,0,1,0},
		{0,0,0,0,0,0,1},
	}))
}
func findNumberOfTruckPopups(row, col int, data [][]int) int {
	marked := make([][]int, row)
	for i := range marked {
		marked[i] = make([]int, col)
	}

	count := 0
	for i, r := range data {
		for j, c := range r {
			if c == 0 || marked[i][j] == 1 {
				continue
			}
			n := node{x: i, y: j}
			marked[i][j] = 1
			count++

			area := n.Neighbours(row, col)
			for len(area) > 0 {
				newArea := make([]node, 0)
				for _, neighbor := range area {
					if data[neighbor.x][neighbor.y] == 1 && marked[neighbor.x][neighbor.y] != 1 {
						marked[neighbor.x][neighbor.y] = 1
						newArea = append(newArea, node{x: neighbor.x, y: neighbor.y})
					}
				}
				area = newArea
			}
		}
	}
	return count
}
