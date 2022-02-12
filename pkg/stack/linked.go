package stack

// LinkedNode is a record of a linked list.
type LinkedNode struct {
	data interface{}
	next *LinkedNode
}

// LinkedStack a linear collection of data nodes.
type LinkedStack struct {
	head *LinkedNode
}

// NewLinkedStack an empty stack.
func NewLinkedStack() *LinkedStack {
	return &LinkedStack{}
}

// Push data onto stack.
func (s *LinkedStack) Push(data interface{}) {
	node := &LinkedNode{data: data}

	node.next = s.head
	s.head = node
}

// Pop from the stack
func (s *LinkedStack) Pop() interface{} {
	if s.head == nil {
		return nil
	}

	head := s.head
	s.head = s.head.next

	return head.data
}

// Peek from the stack
func (s *LinkedStack) Peek() interface{} {
	if s.head == nil {
		return nil
	}

	return s.head.data
}
