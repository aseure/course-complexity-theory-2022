package tree

import "golang.org/x/exp/constraints"

type AVLTreeNode[T constraints.Ordered] struct {
	value  T
	height int
	left   *AVLTreeNode[T]
	right  *AVLTreeNode[T]
}

func NewAVLTreeNode[T constraints.Ordered](value T) *AVLTreeNode[T] {
	return &AVLTreeNode[T]{
		value:  value,
		height: 1,
		left:   nil,
		right:  nil,
	}
}

type AVLTreeNodeVisitor[T constraints.Ordered] func(n *AVLTreeNode[T])

func (n *AVLTreeNode[T]) Accept(v AVLTreeNodeVisitor[T]) {
	if n == nil {
		return
	}
	n.left.Accept(v)
	v(n)
	n.right.Accept(v)
}

func (n *AVLTreeNode[T]) Size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.Size() + n.right.Size()
}

func (n *AVLTreeNode[T]) Height() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AVLTreeNode[T]) Balance() int {
	if n == nil {
		return 0
	}
	return n.left.Height() - n.right.Height()
}

func (n *AVLTreeNode[T]) Insert(value T) *AVLTreeNode[T] {
	if n == nil {
		return NewAVLTreeNode[T](value)
	}
	if value < n.value {
		n.left = n.left.Insert(value)
	}
	if value > n.value {
		n.right = n.right.Insert(value)
	}

	n.height = 1 + max(n.left.Height(), n.right.Height())
	b := n.Balance()

	if b > 1 && value < n.left.value {
		return n.rightRotate()
	}

	if b > 1 && value > n.left.value {
		n.left = n.left.leftRotate()
		return n.rightRotate()
	}

	if b < -1 && value > n.right.value {
		return n.leftRotate()
	}

	if b < -1 && value < n.right.value {
		n.right = n.right.rightRotate()
		return n.leftRotate()
	}

	return n
}

func (n *AVLTreeNode[T]) rightRotate() *AVLTreeNode[T] {
	oldLeft := n.left
	oldLeftRight := n.left.right

	n.left = oldLeftRight
	oldLeft.right = n

	n.height = 1 + max(n.left.Height(), n.right.Height())
	oldLeft.height = 1 + max(oldLeft.left.Height(), oldLeft.right.Height())

	return oldLeft
}

func (n *AVLTreeNode[T]) leftRotate() *AVLTreeNode[T] {
	oldRight := n.right
	oldRightLeft := n.right.left

	n.right = oldRightLeft
	oldRight.left = n

	n.height = 1 + max(n.left.Height(), n.right.Height())
	oldRight.height = 1 + max(oldRight.left.Height(), oldRight.right.Height())

	return oldRight
}

func (n *AVLTreeNode[T]) Remove(value T) (*AVLTreeNode[T], bool) {
	if n == nil {
		return nil, false
	}

	var ok bool

	if value < n.value {
		n.left, ok = n.left.Remove(value)
	} else if value > n.value {
		n.right, ok = n.right.Remove(value)
	} else {
		if n.left != nil {
			n.value = n.left.getMax()
			n.left, ok = n.left.Remove(n.value)
		} else if n.right != nil {
			n.value = n.right.getMin()
			n.right, ok = n.right.Remove(n.value)
		} else {
			return nil, true
		}
	}

	n.height = 1 + max(n.left.Height(), n.right.Height())
	b := n.Balance()

	if b > 1 && n.left.Balance() >= 0 {
		return n.rightRotate(), ok
	}

	if b > 1 && n.left.Balance() < 0 {
		n.left = n.left.leftRotate()
		return n.rightRotate(), ok
	}

	if b < -1 && n.right.Balance() <= 0 {
		return n.leftRotate(), ok
	}

	if b < -1 && n.right.Balance() > 0 {
		n.right = n.right.rightRotate()
		return n.leftRotate(), ok
	}

	return n, ok
}

func (n *AVLTreeNode[T]) Search(value T) bool {
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

func (n *AVLTreeNode[T]) getMin() T {
	if n.left == nil {
		return n.value
	}
	return n.left.getMin()
}

func (n *AVLTreeNode[T]) getMax() T {
	if n.right == nil {
		return n.value
	}
	return n.right.getMax()
}
