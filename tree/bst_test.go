package tree

import (
	"fmt"
	"testing"
)

func TestNode_PrintInOrder(t *testing.T) {
	root := NewTree()
	fmt.Println(root.InOrder())
}
