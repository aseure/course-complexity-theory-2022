package tree

import "golang.org/x/exp/constraints"

type BinaryTree[T constraints.Ordered] interface {
	Size() int
	Height() int
	Insert(value T)
	Remove(value T) bool
	Search(value T) bool
	Display() error
}
