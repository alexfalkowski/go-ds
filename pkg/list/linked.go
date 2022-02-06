package list

// LinkedNode is a record of a linked list.
type LinkedNode struct {
	Data interface{}
	next *LinkedNode
}

// LinkedList a linear collection of data elements whose order is not given by their physical placement in memory.
type LinkedList struct {
	head   *LinkedNode
	tail   *LinkedNode
	length uint64
}

// NewLinkedList an empty linked list.
func NewLinkedList() *LinkedList {
	return &LinkedList{}
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
		l.tail = node
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

// Length of the Length
func (l *LinkedList) Length() uint64 {
	return l.length
}
