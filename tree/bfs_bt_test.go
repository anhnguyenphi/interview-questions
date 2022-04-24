package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func enqueue(queue []*Node, n *Node) []*Node {
	queue = append(queue, n)
	return queue
}

func dequeue(queue []*Node) ([]*Node, *Node) {
	return queue[1:], queue[0]
}

func bfs(root *Node) []int {
	queue := make([]*Node, 0)
	queue = enqueue(queue, root)
	res := make([]int, 0)
	for {
		if len(queue) == 0 {
			break
		}
		newQueue, head := dequeue(queue)
		queue = newQueue
		if head.Left != nil {
			queue = enqueue(queue, head.Left)
		}
		if head.Right != nil {
			queue = enqueue(queue, head.Right)
		}
		res = append(res, head.Val)
	}
	return res
}

func TestBFS(t *testing.T) {
	tree := NewTree()
	res := bfs(tree)
	assert.Equal(t, res, []int{4, 3, 5, 2, 6, 1})
}
