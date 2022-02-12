package stack

// Stack operations
type Stack interface {
	Push(data interface{})
	Pop() interface{}
	Peek() interface{}
}
