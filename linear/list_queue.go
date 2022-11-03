package linear

type ListQueue[T any] struct {
	list *DoublyLinkedList[T]
}

func NewListQueue[T any]() *ListQueue[T] {
	return &ListQueue[T]{
		list: NewDoublyLinkedList[T](),
	}
}

func (q *ListQueue[T]) IsEmpty() bool {
	return q.list.Size() == 0
}

func (q *ListQueue[T]) Push(value T) bool {
	return q.list.PushBack(value)
}

func (q *ListQueue[T]) Pop() (T, bool) {
	return q.list.DeleteAt(0)
}
