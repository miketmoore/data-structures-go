package queue

import (
	"github.com/miketmoore/data-structures-go/linkedlist"
)

// Queue represents the FIFO (first in, first out) principle
type Queue struct {
	list linkedlist.LinkedList
}

// New returns a new queue
func New() Queue {
	return Queue{list: linkedlist.New()}
}

// Enqueue adds a new node to the end of the list
func (q *Queue) Enqueue(node *linkedlist.Node) {
	q.list.Add(node)
}

// Dequeue removes the first node (head) from the list
func (q *Queue) Dequeue() *linkedlist.Node {
	return q.list.RemoveHead()
}

// Peek returns but does not remove the first node (head) in the list
func (q *Queue) Peek() *linkedlist.Node {
	return q.list.Tail
}

// IsEmpty indicates if the list is empty or not
func (q *Queue) IsEmpty() bool {
	return q.list.Size() == 0
}
