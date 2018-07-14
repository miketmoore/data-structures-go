package stack_test

import (
	"testing"

	"github.com/miketmoore/data-structures-go/linkedlist"
	"github.com/miketmoore/data-structures-go/stack"
)

func TestNew(t *testing.T) {
	s := stack.New()
	expected := stack.Stack{}
	if s != expected {
		t.Fatal("New failed")
	}
}

func TestIntegration(t *testing.T) {
	s := stack.New()

	ok(t, s.Peek() == nil)

	s.Push(&linkedlist.Node{Value: 1})
	ok(t, !s.IsEmpty())

	ok(t, s.Peek().Value == 1)

	node := s.Pop()
	ok(t, node.Value == 1)
	ok(t, s.IsEmpty())

	node = s.Peek()
	ok(t, node == nil)

	s.Push(&linkedlist.Node{Value: 100})
	ok(t, s.Peek().Value == 100)
	ok(t, !s.IsEmpty())

	s.Push(&linkedlist.Node{Value: 200})
	ok(t, s.Peek().Value == 200)
	ok(t, !s.IsEmpty())

	node = s.Pop()
	ok(t, node.Value == 200)
	ok(t, !s.IsEmpty())

	node = s.Pop()
	ok(t, node.Value == 100)
	ok(t, s.IsEmpty())
}

func ok(t *testing.T, v bool) {
	if v == false {
		t.Fatal("not ok")
	}
}
