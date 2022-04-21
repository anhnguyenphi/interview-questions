package tree

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

// link: https://www.geeksforgeeks.org/sorted-array-to-balanced-bst/
// Time: O(N), Space: O(1)
func sortedArrToBst(arr []int) *Node {
	arrLen := len(arr)
	if arrLen == 1 {
		return &Node{Val: arr[0]}
	}
	mid := arrLen / 2
	root := &Node{Val: arr[mid]}
	root.Left = sortedArrToBst(arr[:mid])
	if mid+1 < arrLen {
		root.Right = sortedArrToBst(arr[mid+1:])
	}
	return root
}

func TestSortedArrToBST(t *testing.T) {
	root := sortedArrToBst([]int{1, 2, 3, 4, 5, 6, 7})
	assert.Equal(t, root.PreOrder(), []int{4, 2, 1, 3, 6, 5, 7})
}
