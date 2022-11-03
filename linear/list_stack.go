package linear

type ListStack[T any] struct {
	list *DoublyLinkedList[T]
}

func NewListStack[T any]() *ListStack[T] {
	return &ListStack[T]{
		list: NewDoublyLinkedList[T](),
	}
}

func (s *ListStack[T]) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s *ListStack[T]) Push(value T) bool {
	return s.list.PushFront(value)
}

func (s *ListStack[T]) Pop() (T, bool) {
	return s.list.DeleteAt(0)
}
