package linear

type DoublyLinkedList[T any] struct {
	head *DoublyLinkedListNode[T]
	tail *DoublyLinkedListNode[T]
}

type DoublyLinkedListNode[T any] struct {
	value T
	prev  *DoublyLinkedListNode[T]
	next  *DoublyLinkedListNode[T]
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head: nil,
		tail: nil,
	}
}

func (l *DoublyLinkedList[T]) Size() int {
	size := 0
	node := l.head
	for node != nil {
		size++
		node = node.next
	}
	return size
}

func (l *DoublyLinkedList[T]) Values() []T {
	var values []T
	node := l.head
	for node != nil {
		values = append(values, node.value)
		node = node.next
	}
	return values
}

func (l *DoublyLinkedList[T]) GetAt(position int) (T, bool) {
	node := l.head
	for i := 0; i < position && node != nil; i++ {
		node = node.next
	}
	if node == nil {
		return *new(T), false
	}
	return node.value, true
}

func (l *DoublyLinkedList[T]) InsertAt(position int, value T) bool {
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
	newNode := &DoublyLinkedListNode[T]{
		value: value,
		prev:  prev,
		next:  prev.next,
	}
	if prev.next != nil {
		prev.next.prev = newNode
	}
	prev.next = newNode
	return true
}

func (l *DoublyLinkedList[T]) DeleteAt(position int) (T, bool) {
	if l.head == nil {
		return *new(T), false
	}
	if position == 0 {
		value := l.head.value
		if l.head.next == nil {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
			l.head.prev = nil
		}
		return value, true
	}
	i, prev := 1, l.head
	for i < position && prev != nil {
		prev = prev.next
		i++
	}
	if prev == nil || prev.next == nil {
		return *new(T), false
	}
	value := prev.next.value
	if prev.next.next != nil {
		prev.next.next.prev = prev
	}
	prev.next = prev.next.next
	return value, true
}

func (l *DoublyLinkedList[T]) PushBack(value T) bool {
	newNode := &DoublyLinkedListNode[T]{
		value: value,
		prev:  l.tail,
		next:  nil,
	}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
		return true
	}
	l.tail.next = newNode
	l.tail = newNode
	return true
}

func (l *DoublyLinkedList[T]) PushFront(value T) bool {
	newNode := &DoublyLinkedListNode[T]{
		value: value,
		prev:  nil,
		next:  l.head,
	}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return true
	}
	l.head.prev = newNode
	l.head = newNode
	return true
}
