package linear

type List[T any] interface {
	Size() T
	Values() []T
	GetAt(position T) (T, bool)
	InsertAt(position, value T) bool
	DeleteAt(position T) (T, bool)
	PushBack(value T) bool
	PushFront(value T) bool
}
