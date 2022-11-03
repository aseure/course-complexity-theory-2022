package linear

type ArrayStack[T any] struct {
	list *ArrayList[T]
}

func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{
		list: NewArrayList[T](),
	}
}

func (s *ArrayStack[T]) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s *ArrayStack[T]) Push(value T) bool {
	return s.list.PushBack(value)
}

func (s *ArrayStack[T]) Pop() (T, bool) {
	return s.list.DeleteAt(s.list.Size() - 1)
}
