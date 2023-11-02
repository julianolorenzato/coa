package tree

type node struct {
	Val   any
	Left  *node
	Right *node
}

func newNode(val any) *node {
	return &node{Val: val, Left: nil, Right: nil}
}

type BinaryTree struct {
	Root *node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

func (t *BinaryTree) PreOrder(cb func(*node)) {
	var preOrder func(n *node)

	preOrder = func(n *node) {
		if n == nil {
			return
		}

		cb(n)
		preOrder(n.Left)
		preOrder(n.Right)
	}

	preOrder(t.Root)
}

func (t *BinaryTree) InOrder(cb func(*node)) {
	var inOrder func(n *node)

	inOrder = func(n *node) {
		if n == nil {
			return
		}

		inOrder(n.Left)
		cb(n)
		inOrder(n.Right)
	}

	inOrder(t.Root)
}

func (t *BinaryTree) PostOrder(cb func(*node)) {
	var postOrder func(n *node)

	postOrder = func(n *node) {
		if n == nil {
			return
		}

		postOrder(n.Left)
		postOrder(n.Right)
		cb(n)
	}

	postOrder(t.Root)
}

func NewCompleteBinaryTree(arr []any) *BinaryTree {
	bt := NewBinaryTree()

	if len(arr) == 0 {
		return bt
	}

	bt.Root = mountTree(arr, 0)

	return bt
}

func mountTree(arr []any, i int) *node {
	var nd *node = nil

	if i < len(arr) {
		nd = newNode(arr[i])
		nd.Left = mountTree(arr, 2*i+1)
		nd.Right = mountTree(arr, 2*i+2)
	}

	return nd
}
