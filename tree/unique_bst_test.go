package tree

import (
	"fmt"
	"testing"
)

// link: https://leetcode.com/problems/unique-binary-search-trees-ii/
func generateTrees(n int) []*Node {
	res := make([]*Node, 0)
	res = append(res, &Node{Val: 1})
	for i := 2; i <= n; i++ {
		newRes := make([]*Node, 0)
		for _, node := range res {
			newRes = append(newRes, addNode(node, i)...)
		}
		res = newRes
	}
	return res
}

func addNode(root *Node, val int) []*Node {
	if root == nil {
		return []*Node{
			{
				Val: val,
			},
		}
	}

	res := make([]*Node, 0)
	newNode := &Node{Val: val}
	if root.Val < val {
		newNode.Left = root
		res = append(res, newNode)
		for _, n := range addNode(root.Right, val) {
			res = append(res, &Node{
				Val:   root.Val,
				Left:  root.Left,
				Right: n,
			})
		}
	} else {
		newNode.Right = root
		res = append(res, newNode)
		for _, n := range addNode(root.Left, val) {
			res = append(res, &Node{
				Val:   root.Val,
				Left:  n,
				Right: root.Right,
			})
		}
	}
	return res
}

func TestGenerateTrees(t *testing.T) {
	res := generateTrees(3)
	for _, tree := range res {
		fmt.Println(tree.PreOrder())
	}
}
