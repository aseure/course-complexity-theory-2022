package tree

import "golang.org/x/exp/constraints"

type BinarySearchTreeNode[T constraints.Ordered] struct {
	value T
	left  *BinarySearchTreeNode[T]
	right *BinarySearchTreeNode[T]
}

func NewBinarySearchTreeNode[T constraints.Ordered](value T) *BinarySearchTreeNode[T] {
	return &BinarySearchTreeNode[T]{
		value: value,
		left:  nil,
		right: nil,
	}
}

type BinarySearchTreeNodeVisitor[T constraints.Ordered] func(n *BinarySearchTreeNode[T])

func (n *BinarySearchTreeNode[T]) Accept(v BinarySearchTreeNodeVisitor[T]) {
	if n == nil {
		return
	}
	n.left.Accept(v)
	v(n)
	n.right.Accept(v)
}

func (n *BinarySearchTreeNode[T]) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.Size() + n.right.Size()
}

func (n *BinarySearchTreeNode[T]) Height() int {
	if n == nil {
		return 0
	}
	return 1 + max(n.left.Height(), n.right.Height())
}

func (n *BinarySearchTreeNode[T]) Insert(value T) *BinarySearchTreeNode[T] {
	if n == nil {
		return NewBinarySearchTreeNode(value)
	}
	if value < n.value {
		n.left = n.left.Insert(value)
	}
	if value > n.value {
		n.right = n.right.Insert(value)
	}
	return n
}

func (n *BinarySearchTreeNode[T]) Remove(value T) (*BinarySearchTreeNode[T], bool) {
	if n == nil {
		return nil, false
	}

	var ok bool

	if value < n.value {
		n.left, ok = n.left.Remove(value)
		return n, ok
	}

	if value > n.value {
		n.right, ok = n.right.Remove(value)
		return n, ok
	}

	if n.left != nil {
		n.value = n.left.getMax()
		n.left, ok = n.left.Remove(n.value)
		return n, ok
	}

	if n.right != nil {
		n.value = n.right.getMin()
		n.right, ok = n.right.Remove(n.value)
		return n, ok
	}

	return nil, true
}

func (n *BinarySearchTreeNode[T]) Search(value T) bool {
	if n == nil {
		return false
	}
	if value < n.value {
		return n.left.Search(value)
	}
	if value > n.value {
		return n.right.Search(value)
	}
	return true
}

func (n *BinarySearchTreeNode[T]) getMin() T {
	if n.left == nil {
		return n.value
	}
	return n.left.getMin()
}

func (n *BinarySearchTreeNode[T]) getMax() T {
	if n.right == nil {
		return n.value
	}
	return n.right.getMax()
}
