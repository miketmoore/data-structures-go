package queue_test

import (
	"testing"

	"github.com/miketmoore/data-structures-go/linkedlist"
	"github.com/miketmoore/data-structures-go/queue"
)

func TestNew(t *testing.T) {
	q := queue.New()
	expected := queue.Queue{}
	if q != expected {
		t.Fatal("new failed")
	}
}

func TestIntegration(t *testing.T) {
	q := queue.New()
	ok(t, q.IsEmpty())

	ok(t, q.Peek() == nil)

	q.Enqueue(&linkedlist.Node{Value: 1})
	ok(t, q.Peek().Value == 1)
	ok(t, !q.IsEmpty())

	node := q.Dequeue()
	ok(t, node.Value == 1)
	ok(t, q.Peek() == nil)
	ok(t, q.IsEmpty())

	q.Enqueue(&linkedlist.Node{Value: 100})
	ok(t, q.Peek().Value == 100)
	ok(t, !q.IsEmpty())

	q.Enqueue(&linkedlist.Node{Value: 200})
	ok(t, q.Peek().Value == 200)
	ok(t, !q.IsEmpty())

	node = q.Dequeue()
	ok(t, node.Value == 100)
	ok(t, q.Peek().Value == 200)
	ok(t, !q.IsEmpty())

	node = q.Dequeue()
	ok(t, node.Value == 200)
	ok(t, q.Peek() == nil)
	ok(t, q.IsEmpty())
}

func TestIntegrationNames(t *testing.T) {
	q := queue.New()

	q.Enqueue(&linkedlist.Node{Value: "Mike"})
	q.Enqueue(&linkedlist.Node{Value: "Tarzan"})
	q.Enqueue(&linkedlist.Node{Value: "Jane"})
	q.Enqueue(&linkedlist.Node{Value: "Penny"})

	ok(t, q.Dequeue().Value == "Mike")
	ok(t, q.Dequeue().Value == "Tarzan")
	ok(t, q.Dequeue().Value == "Jane")
	ok(t, q.Dequeue().Value == "Penny")
}

func ok(t *testing.T, v bool) {
	if v == false {
		t.Fatal("not ok")
	}
}
