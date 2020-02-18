package topological_sorting

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// link: https://www.geeksforgeeks.org/topological-sorting/

type Graph struct {
	V int
	Adj [][]int
}

func NewGraph(v int) *Graph {
	return &Graph{
		V: v,
		Adj: make([][]int, v),
	}
}

func (g *Graph) AddEdge(v int, ws... int) {
	for _, w := range ws {
		g.Adj[v] = append(g.Adj[v], w)
	}
}

func (g *Graph) topologicalSortUtil(v int, stack *[]int, visited *[]bool) {
	if (*visited)[v] {
		return
	}

	(*visited)[v] = true

	for i := 0; i < len(g.Adj[v]); i++ {
		if !(*visited)[g.Adj[v][i]] {
			g.topologicalSortUtil(g.Adj[v][i], stack, visited)
		}
	}

	*stack = append(*stack, v)
}

func (g *Graph) TopologicalSort() []int {
	visited := make([]bool, g.V)
	stack := make([]int, 0)
	for i := 0 ; i < g.V; i++ {
		if !visited[i] {
			g.topologicalSortUtil(i, &stack, &visited)
		}
	}

	result := make([]int, len(stack))
	for idx, ele := range stack {
		result[len(stack) - idx - 1] = ele
	}
	return result
}

func TestCase1(t *testing.T) {
	g := NewGraph(6)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)
	g.AddEdge(4, 0, 1)
	g.AddEdge(5, 2, 0)

	assert.Equal(t, g.TopologicalSort(), []int{5, 4, 2, 3, 1, 0})
}