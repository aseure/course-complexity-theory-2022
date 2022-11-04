package tree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MaxHeap[T constraints.Ordered] struct {
	nodes []T
}

func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{
		nodes: nil,
	}
}

func (h *MaxHeap[T]) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *MaxHeap[T]) String() string {
	return fmt.Sprintf("%v", h.nodes)
}

func (h *MaxHeap[T]) Size() int {
	return len(h.nodes)
}

func (h *MaxHeap[T]) Insert(value T) {
	h.nodes = append(h.nodes, value)
	h.heapifyUp(len(h.nodes) - 1)
}

func (h *MaxHeap[T]) RemoveMax() (T, bool) {
	if len(h.nodes) == 0 {
		return *new(T), false
	}
	max := h.nodes[0]
	h.nodes[0] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.heapifyDown(0)
	return max, true
}

func (h *MaxHeap[T]) Display() error {
	start := func(w func(format string, a ...interface{})) {
		visitor := func(i int) {
			_, leftChild, rightChild := getRelationshipIndices(i)

			if leftChild < len(h.nodes) {
				w("  %d -> %d;\n", h.nodes[i], h.nodes[leftChild])
			}
			if rightChild < len(h.nodes) {
				w("  %d -> %d;\n", h.nodes[i], h.nodes[rightChild])
			}
		}

		if len(h.nodes) > 0 {
			if len(h.nodes) == 1 {
				w("  %d;\n", h.nodes[0])
			} else {
				h.Accept(0, visitor)
			}
		}
	}

	return display(start)
}

func (h *MaxHeap[T]) Accept(i int, visitor func(int)) {
	if i >= len(h.nodes) {
		return
	}
	_, leftChild, rightChild := getRelationshipIndices(i)
	h.Accept(leftChild, visitor)
	visitor(i)
	h.Accept(rightChild, visitor)
}

func (h *MaxHeap[T]) heapifyUp(i int) {
	for i > 0 {
		parent, _, _ := getRelationshipIndices(i)
		isParentSmaller := h.nodes[parent] < h.nodes[i]
		if isParentSmaller {
			h.swap(parent, i)
		}
		i = parent
	}
}

func (h *MaxHeap[T]) heapifyDown(i int) {
	for {
		_, leftChild, rightChild := getRelationshipIndices(i)
		isLeftChildBigger := leftChild < len(h.nodes) && h.nodes[i] < h.nodes[leftChild]
		isRightChildBigger := rightChild < len(h.nodes) && h.nodes[i] < h.nodes[rightChild]

		if isLeftChildBigger && isRightChildBigger {
			if h.nodes[leftChild] < h.nodes[rightChild] {
				h.swap(i, rightChild)
				i = rightChild
			} else {
				h.swap(i, leftChild)
				i = leftChild
			}
		} else if isLeftChildBigger {
			h.swap(i, leftChild)
			i = leftChild
		} else if isRightChildBigger {
			h.swap(i, rightChild)
			i = rightChild
		} else {
			return
		}
	}
}

func getRelationshipIndices(i int) (parent, leftChild, rightChild int) {
	return (i - 1) / 2, 2*i + 1, 2*i + 2
}
