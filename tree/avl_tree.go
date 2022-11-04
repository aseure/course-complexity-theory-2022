package tree

import "golang.org/x/exp/constraints"

type AVLTree[T constraints.Ordered] struct {
	root *AVLTreeNode[T]
}

func NewAVLTree[T constraints.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{
		root: nil,
	}
}

func (t *AVLTree[T]) Size() int {
	return t.root.Size()
}

func (t *AVLTree[T]) Height() int {
	return t.root.Height()
}

func (t *AVLTree[T]) Insert(value T) {
	t.root = t.root.Insert(value)
}

func (t *AVLTree[T]) Remove(value T) bool {
	var ok bool
	t.root, ok = t.root.Remove(value)
	return ok
}

func (t *AVLTree[T]) Search(value T) bool {
	return t.root.Search(value)
}

func (t *AVLTree[T]) Display() error {
	start := func(w func(format string, a ...interface{})) {
		visitor := func(n *AVLTreeNode[T]) {
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

func (t *AVLTree[T]) DepthFirstTraversal() (values []T) {
	if t.root == nil {
		return
	}
	t.root.Accept(func(n *AVLTreeNode[T]) {
		values = append(values, n.value)
	})
	return
}

func (t *AVLTree[T]) BreadthFirstTraversal() (values []T) {
	if t.root == nil {
		return
	}

	queue := []*AVLTreeNode[T]{t.root}

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n != nil {
			queue = append(queue, n.left, n.right)
			values = append(values, n.value)
		}
	}

	return
}
