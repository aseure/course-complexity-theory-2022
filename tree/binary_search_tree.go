package tree

import "golang.org/x/exp/constraints"

type BinarySearchTree[T constraints.Ordered] struct {
	root *BinarySearchTreeNode[T]
}

func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		root: nil,
	}
}

func (t *BinarySearchTree[T]) Size() int {
	return t.root.Size()
}

func (t *BinarySearchTree[T]) Height() int {
	return t.root.Height()
}

func (t *BinarySearchTree[T]) Insert(value T) {
	t.root = t.root.Insert(value)
}

func (t *BinarySearchTree[T]) Remove(value T) bool {
	var ok bool
	t.root, ok = t.root.Remove(value)
	return ok
}

func (t *BinarySearchTree[T]) Search(value T) bool {
	return t.root.Search(value)
}

func (t *BinarySearchTree[T]) Display() error {
	start := func(w func(format string, a ...interface{})) {
		visitor := func(n *BinarySearchTreeNode[T]) {
			if n.left != nil {
				w("  %d -> %d;\n", n.value, n.left.value)
			}
			if n.right != nil {
				w("  %d -> %d;\n", n.value, n.right.value)
			}
		}

		if t.root != nil {
			if t.root.left == nil && t.root.right == nil {
				w("  %d;\n", t.root.value)
			} else {
				t.root.Accept(visitor)
			}
		}
	}

	return display(start)
}
