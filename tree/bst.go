package tree

type (
	Node struct {
		Val   int
		Left  *Node
		Right *Node
	}
)

// will return sorted order
func (n *Node) InOrder() []int {
	var inOrder []int
	if n.Left != nil {
		inOrder = n.Left.InOrder()
	}
	inOrder = append(inOrder, n.Val)
	if n.Right != nil {
		inOrder = append(inOrder, n.Right.InOrder()...)

	}
	return inOrder
}

// will return sorted order
func (n *Node) PreOrder() []int {
	var inOrder []int
	inOrder = append(inOrder, n.Val)
	if n.Left != nil {
		inOrder = append(inOrder, n.Left.PreOrder()...)
	}
	if n.Right != nil {
		inOrder = append(inOrder, n.Right.PreOrder()...)

	}
	return inOrder
}

func insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{Val: val}
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root
}

func NewTree() *Node {
	root := insert(nil, 4)
	insert(root, 3)
	insert(root, 2)
	insert(root, 1)
	insert(root, 5)
	insert(root, 6)
	return root
}
