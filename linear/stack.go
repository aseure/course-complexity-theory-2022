package linear

type Stack[T any] interface {
	IsEmpty() bool
	Push(value T) bool
	Pop() (T, bool)
}
