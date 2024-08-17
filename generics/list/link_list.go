package list

type linkNode[T comparable] struct {
	data T
	next *linkNode[T]
}

type LinkList[T comparable] struct {
	head   *linkNode[T]
	length int
}

// NewLinkList creates a new empty list with head node.
func NewLinkList[T comparable]() *LinkList[T] {
	head := &linkNode[T]{}
	return &LinkList[T]{head: head}
}

// InsertFirst inserts a new element at the beginning of the list.
func (list *LinkList[T]) InsertFirst(v T) {
	data := &linkNode[T]{data: v}
	if list.head != nil {
		data.next = list.head
	}
	list.head = data
}

// InsertLast inserts a new element at the end of the list.
func (list *LinkList[T]) InsertLast(v T) {
	data := &linkNode[T]{data: v}
	if list.head == nil {
		list.head = data
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = data
}

// RemoveByValue removes an element from the list by its value.
func (list *LinkList[T]) RemoveByValue(v T) bool {
	if list.head == nil {
		return false
	}
	if equal(list.head.data, v) {
		list.head = list.head.next
		return true
	} else {
		current := list.head
		for current.next != nil {
			if current.next.data == v {
				current.next = current.next.next
				return true
			}
			current = current.next
		}
		return false
	}
}

func equal[T comparable](data, v T) bool {
	if data == v {
		return true
	}
	return false
}

// RemoveByIndex removes an element from the list by its index.
func (list *LinkList[T]) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head = list.head.next
		return true
	}

	current := list.head
	for u := 1; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	return false
}

// SearchValue searches an element in the list by its value.
// if found, return true, otherwise return false.
func (list *LinkList[T]) SearchValue(v T) bool {
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

// GetFirst gets the first element of the list.
// if list is empty, return false and zero value.
// if list is not empty, return true and the first element.
func (list *LinkList[T]) GetFirst() (T, bool) {
	var v T
	if list.head == nil {
		return v, false
	}
	return list.head.data, true
}

// GetLast gets the last element of the list.
// if list is empty, return false and zero value.
// if list is not empty, return true and the last element.
func (list *LinkList[T]) GetLast() (T, bool) {
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

// GetSize gets the size of the list.
func (list *LinkList[T]) GetSize() int {
	size := 0
	current := list.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

// GetItems gets all the elements of the list.
func (list *LinkList[T]) GetItems() []T {
	var items []T
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}
