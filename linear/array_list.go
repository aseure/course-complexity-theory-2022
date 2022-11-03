package linear

type ArrayList[T any] struct {
	arr      []T
	capacity int
	end      int
}

func NewArrayList[T any]() *ArrayList[T] {
	capacity := 10
	return &ArrayList[T]{
		arr:      make([]T, capacity),
		capacity: capacity,
		end:      0,
	}
}

func (l *ArrayList[T]) Size() int {
	return l.end
}

func (l *ArrayList[T]) Values() []T {
	return l.arr[:l.end]
}

func (l *ArrayList[T]) GetAt(position int) (T, bool) {
	if position >= l.end {
		return *new(T), false
	}
	return l.arr[position], true
}

func (l *ArrayList[T]) InsertAt(position int, value T) bool {
	if position < 0 || l.end < position {
		return false
	}
	if l.end == l.capacity {
		l.growCapacity()
	}
	for i := l.end; i > position; i-- {
		l.arr[i] = l.arr[i-1]
	}
	l.arr[position] = value
	l.end++
	return true
}

func (l *ArrayList[T]) DeleteAt(position int) (T, bool) {
	if position < 0 || l.end < position {
		return *new(T), false
	}
	value := l.arr[position]
	for i := position; i < l.end-1; i++ {
		l.arr[i] = l.arr[i+1]
	}
	l.end--
	return value, true
}

func (l *ArrayList[T]) PushBack(value T) bool {
	if l.end == l.capacity {
		l.growCapacity()
	}
	l.arr[l.end] = value
	l.end++
	return true
}

func (l *ArrayList[T]) PushFront(value T) bool {
	if l.end == l.capacity {
		l.growCapacity()
	}
	for i := l.end; i > 0; i-- {
		l.arr[i] = l.arr[i-1]
	}
	l.arr[0] = value
	l.end++
	return true
}

func (l *ArrayList[T]) growCapacity() {
	l.capacity = l.capacity * 2
	newArr := make([]T, l.capacity)
	copy(newArr, l.arr)
	l.arr = newArr
}
