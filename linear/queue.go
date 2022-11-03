package linear

type Queue[T any] interface {
	IsEmpty() bool
	Push(value T) bool
	Pop() (T, bool)
}
