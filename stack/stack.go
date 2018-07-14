package stack

import "github.com/miketmoore/data-structures-go/linkedlist"

// Stack is an adapter on top of LinkedList
// It enforces the last-in first-out (LIFO) principle.
type Stack struct {
	list linkedlist.LinkedList
}

// Push adds the node to the beginning of the list (last in)
func (s *Stack) Push(node *linkedlist.Node) {
	s.list.AddToStart(node)
}

// Pop removes the node at the beginning of the list (first out)
func (s *Stack) Pop() *linkedlist.Node {
	return s.list.RemoveHead()
}

// IsEmpty indicates if the stack is empty or not
func (s *Stack) IsEmpty() bool {
	return s.list.Size() == 0
}

// Peek returns the top node, but does not remove it
func (s *Stack) Peek() *linkedlist.Node {
	return s.list.Head
}

// New returns a new Stack instance
func New() Stack {
	return Stack{list: linkedlist.New()}
}
