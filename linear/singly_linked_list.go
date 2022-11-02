package linear

type SinglyLinkedList[T any] struct {
	head *SinglyLinkedListNode[T]
}

type SinglyLinkedListNode[T any] struct {
	value T
	next  *SinglyLinkedListNode[T]
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: nil,
	}
}

func (l *SinglyLinkedList[T]) Size() int {
	size := 0
	node := l.head
	for node != nil {
		size++
		node = node.next
	}
	return size
}

func (l *SinglyLinkedList[T]) Values() []T {
	var values []T
	node := l.head
	for node != nil {
		values = append(values, node.value)
		node = node.next
	}
	return values
}

func (l *SinglyLinkedList[T]) GetAt(position int) (T, bool) {
	node := l.head
	for i := 0; i < position && node != nil; i++ {
		node = node.next
	}
	if node == nil {
		return *new(T), false
	}
	return node.value, true
}

func (l *SinglyLinkedList[T]) InsertAt(position int, value T) bool {
	if position == 0 {
		return l.PushFront(value)
	}
	if l.head == nil {
		return false
	}
	prev := l.head
	for i := 1; i < position && prev != nil; i++ {
		prev = prev.next
	}
	if prev == nil {
		return false
	}
	newNode := &SinglyLinkedListNode[T]{
		value: value,
		next:  prev.next,
	}
	prev.next = newNode
	return true
}

func (l *SinglyLinkedList[T]) DeleteAt(position int) (T, bool) {
	if l.head == nil || position < 0 {
		return *new(T), false
	}
	if position == 0 {
		value := l.head.value
		l.head = l.head.next
		return value, true
	}
	prev := l.head
	for i := 1; i < position && prev != nil; i++ {
		prev = prev.next
	}
	if prev == nil {
		return *new(T), false
	}
	if prev.next == nil {
		return *new(T), false
	}
	value := prev.next.value
	prev.next = prev.next.next
	return value, true
}

func (l *SinglyLinkedList[T]) PushBack(value T) bool {
	newNode := &SinglyLinkedListNode[T]{
		value: value,
		next:  nil,
	}
	if l.head == nil {
		l.head = newNode
		return true
	}
	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = newNode
	return true
}

func (l *SinglyLinkedList[T]) PushFront(value T) bool {
	newNode := &SinglyLinkedListNode[T]{
		value: value,
		next:  l.head,
	}
	l.head = newNode
	return true
}
