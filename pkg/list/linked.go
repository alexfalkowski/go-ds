package list

// LinkedNode is a record of a linked list.
type LinkedNode struct {
	Data interface{}
	next *LinkedNode
}

// LinkedList a linear collection of data nodes.
type LinkedList struct {
	head   *LinkedNode
	length uint64
}

// NewLinkedList an empty linked list.
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Add the data at an index.
func (l *LinkedList) Add(index uint64, data interface{}) *LinkedNode {
	if index > l.length || data == nil {
		return nil
	}

	if index == 0 {
		return l.AddHead(data)
	}

	if index == l.length {
		return l.AddTail(data)
	}

	current := l.head

	var i uint64

	for i = 1; i < index; i++ {
		current = current.next
	}

	l.length++

	node := &LinkedNode{Data: data}

	node.next = current.next
	current.next = node

	return node
}

// AddHead to the linked list.
func (l *LinkedList) AddHead(data interface{}) *LinkedNode {
	if data == nil {
		return nil
	}

	l.length++

	node := &LinkedNode{Data: data}

	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}

	return node
}

// RemoveHead from the linked list.
func (l *LinkedList) RemoveHead() *LinkedNode {
	if l.head == nil {
		return nil
	}

	l.length--

	prevHead := l.head
	head := prevHead.next
	l.head = head

	return prevHead
}

// AddTail to the linked list.
func (l *LinkedList) AddTail(data interface{}) *LinkedNode {
	if data == nil {
		return nil
	}

	l.length++

	node := &LinkedNode{Data: data}

	if l.head == nil {
		l.head = node
	} else {
		tail := l.head

		for tail.next != nil {
			tail = tail.next
		}

		tail.next = node
	}

	return node
}

// RemoveTail from the linked list.
func (l *LinkedList) RemoveTail() *LinkedNode {
	if l.head == nil {
		return nil
	}

	l.length--

	prevTail := l.head
	tail := l.head

	for tail.next != nil {
		prevTail = tail
		tail = tail.next
	}

	prevTail.next = nil

	return tail
}

// Length of the Length
func (l *LinkedList) Length() uint64 {
	return l.length
}
