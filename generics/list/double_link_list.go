package list

type doubleLinkListNode[T comparable] struct {
	data T
	next *doubleLinkListNode[T]
	prev *doubleLinkListNode[T]
}

type DoubleLinkList[T comparable] struct {
	head   *doubleLinkListNode[T]
	tail   *doubleLinkListNode[T]
	length int
}

// NewDoubleLinkList creates a new empty double-current list.
func NewDoubleLinkList[T comparable]() *DoubleLinkList[T] {
	head := &doubleLinkListNode[T]{}
	tail := &doubleLinkListNode[T]{}
	head.next = tail
	tail.prev = head
	return &DoubleLinkList[T]{head: head, tail: tail}
}

// InsertFirst inserts a new element with value v at the front of the list.
func (list *DoubleLinkList[T]) InsertFirst(v T) {
	data := &doubleLinkListNode[T]{data: v}
	if list.head != nil {
		list.head.prev = data
		data.next = list.head
	}
	list.head = data
}

// InsertLast inserts a new element with value v at the end of the list.
func (list *DoubleLinkList[T]) InsertLast(v T) {
	data := &doubleLinkListNode[T]{data: v}
	if list.head == nil {
		list.head = data
		list.tail = data
	}
	if list.tail != nil {
		list.tail.next = data
		data.prev = list.tail
	}
	list.tail = data
}

// RemoveByValue removes the first element with value v from the list.
func (list *DoubleLinkList[T]) RemoveByValue(v T) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == v {
		list.head = list.head.next
		list.head.prev = nil
		return true
	}
	if list.head.data == v {
		list.tail = list.tail.prev
		list.tail.next = nil
		return true
	}
	current := list.head
	for current.next != nil {
		if list.head.data == v {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

// RemoveByIndex removes the i-th element from the list.
func (list *DoubleLinkList[T]) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head.prev = nil
		list.head = list.head.next
		return false
	}
	current := list.head
	for u := 0; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	return true
}

// SearchValue returns true if the list contains an element with value v.
func (list *DoubleLinkList[T]) SearchValue(v T) bool {
	if list.head == nil {
		return false
	}
	current := list.head
	for current != nil {
		if current.data == v {
			return true
		}
		current = current.next
	}
	return false
}

// GetFirst returns the first element of the list.
func (list *DoubleLinkList[T]) GetFirst() (T, bool) {
	var v T
	if list.head == nil {
		return v, false
	}
	return list.head.data, true
}

// GetLast returns the last element of the list.
func (list *DoubleLinkList[T]) GetLast() (T, bool) {
	var v T
	if list.head == nil {
		return v, false
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current.data, true
}

// GetSize returns the number of elements in the list.
func (list *DoubleLinkList[T]) GetSize() int {
	size := 0
	current := list.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

// GetItemsFromStart returns all elements of the list from the start.
func (list *DoubleLinkList[T]) GetItemsFromStart() (items []T) {
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

// GetItemsFromEnd returns all elements of the list from the end.
func (list *DoubleLinkList[T]) GetItemsFromEnd() (items []T) {
	current := list.tail
	for current != nil {
		items = append(items, current.data)
		current = current.prev
	}
	return items
}
