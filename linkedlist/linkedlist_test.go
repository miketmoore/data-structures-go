package linkedlist_test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/miketmoore/data-structures-go/linkedlist"
)

func TestNewList(t *testing.T) {
	got := newList()
	expected := linkedlist.LinkedList{}
	if got != expected {
		t.Fatal("list is unexpected")
	}
}

func TestAdd(t *testing.T) {
	t.Run("add one", func(t *testing.T) {
		list := newList(1)
		assertListNodes(t, list, 1)
	})
	t.Run("add multiple", func(t *testing.T) {
		list := newList(1, 2, 3)
		assertListNodes(t, list, 1, 2, 3)
	})
}

func TestAddToStart(t *testing.T) {
	list := newList()

	list.AddToStart(&linkedlist.Node{Value: 1})
	assertListNodes(t, list, 1)

	list.AddToStart(&linkedlist.Node{Value: 2})
	assertListNodes(t, list, 2, 1)

	list.AddToStart(&linkedlist.Node{Value: 3})
	assertListNodes(t, list, 3, 2, 1)
}

func TestIterator(t *testing.T) {
	assertions := func(i int, prev, next *linkedlist.Node) {
		if i == 0 {
			if prev != nil {
				t.Fatal("previous is not nil")
			}
		} else {
			if prev == nil {
				t.Fatal("previous is nil")
			}
		}
		if i < 2 {
			if next == nil {
				t.Fatal("next is nil")
			}
		} else {
			if next != nil {
				t.Fatal("next is not nil")
			}
		}
	}
	t.Run("int", func(t *testing.T) {
		list := newList(1, 2, 3)

		it := list.Iterator()

		collect := []int{}
		i := 0
		for it.HasNext() {
			node := it.Next()
			collect = append(collect, node.Value.(int))
			assertions(i, node.Previous(), node.Next())
			i++
		}
		assertSlicesAreEqual(t, []int{1, 2, 3}, collect)
	})
}

func TestDescendingIterator(t *testing.T) {
	list := newList(1, 2, 3)

	it := list.DescendingIterator()

	t.Run("test new", func(t *testing.T) {
		if &it == nil {
			t.Fatal("iterator is nil")
		}
	})

	t.Run("test iteration", func(t *testing.T) {
		collect := []int{}
		for it.HasNext() {
			node := it.Next()
			collect = append(collect, node.Value.(int))
		}
		assertSlicesAreEqual(t, []int{3, 2, 1}, collect)
	})
}

func TestRemoveHead(t *testing.T) {
	t.Run("foo", func(t *testing.T) {
		list := newList()
		ok(t, list.Head == nil)
		list.AddToStart(&linkedlist.Node{Value: 1})
		ok(t, list.Head.Value == 1)
		node := list.RemoveHead()
		ok(t, node.Value == 1)
		ok(t, list.Head == nil)
	})

	list := newList(1, 2, 3)

	t.Run("remove first", func(t *testing.T) {
		node := list.RemoveHead()
		ok(t, node.Value == 1)
		assertListNodes(t, list, 2, 3)
	})

	t.Run("remove second", func(t *testing.T) {
		node := list.RemoveHead()
		ok(t, node.Value == 2)
		assertListNodes(t, list, 3)
	})

	t.Run("remove third", func(t *testing.T) {
		node := list.RemoveHead()
		ok(t, node.Value == 3)
		assertListIsEmpty(t, list)
	})

	t.Run("remove when empty", func(t *testing.T) {
		node := list.RemoveHead()
		ok(t, node == nil)
		assertListIsEmpty(t, list)
	})

}

func TestRemoveTailFail(t *testing.T) {
	list := newList()
	node := list.RemoveTail()
	ok(t, node == nil)
}

func TestRemoveTail(t *testing.T) {
	list := newList(1, 2, 3)

	t.Run("remove 3", func(t *testing.T) {
		node := list.RemoveTail()
		ok(t, node.Value == 3)
		assertListNodes(t, list, 1, 2)
	})

	t.Run("remove 2", func(t *testing.T) {
		node := list.RemoveTail()
		ok(t, node.Value == 2)
		assertListNodes(t, list, 1)
	})

	t.Run("remove 1", func(t *testing.T) {
		node := list.RemoveTail()
		ok(t, node.Value == 1)
		assertListIsEmpty(t, list)
	})

	t.Run("remove when empty", func(t *testing.T) {
		node := list.RemoveTail()
		ok(t, node == nil)
		assertListIsEmpty(t, list)
	})
}

func TestFind(t *testing.T) {
	list := newList(1, 2)

	success, node := list.Find(func(node *linkedlist.Node) bool {
		return node.Value == 1
	})
	ok(t, success)
	assertNodeValue(t, 1, node)

	success, node = list.Find(func(node *linkedlist.Node) bool {
		return node.Value == 2
	})
	ok(t, success)
	assertNodeValue(t, 2, node)

	success, node = list.Find(func(node *linkedlist.Node) bool {
		return node.Value == 100
	})
	ok(t, !success)
	ok(t, node == nil)
}

func TestRemoveFirstOccurrence(t *testing.T) {
	list := newList()

	t.Run("attempt to remove when empty", func(t *testing.T) {
		success := list.RemoveFirstOccurrence(func(node *linkedlist.Node) bool {
			return node.Value == 1
		})
		ok(t, !success)
		assertListIsEmpty(t, list)
	})
	t.Run("remove head", func(t *testing.T) {
		list := newList(1, 2, 3)
		success := list.RemoveFirstOccurrence(func(node *linkedlist.Node) bool {
			return node.Value == 1
		})
		ok(t, success)
		assertListNodes(t, list, 2, 3)
	})
	t.Run("remove middle", func(t *testing.T) {
		// 1-2-3 becomes 1-3
		list := newList(1, 2, 3)
		success := list.RemoveFirstOccurrence(func(node *linkedlist.Node) bool {
			return node.Value == 2
		})
		ok(t, success)
		assertListNodes(t, list, 1, 3)
	})
	t.Run("remove tail", func(t *testing.T) {
		list := newList(1, 2, 3)
		success := list.RemoveFirstOccurrence(func(node *linkedlist.Node) bool {
			return node.Value == 3
		})
		ok(t, success)
		assertListNodes(t, list, 1, 2)
	})
	t.Run("remove no match", func(t *testing.T) {
		list := newList(1, 2, 3)
		success := list.RemoveFirstOccurrence(func(node *linkedlist.Node) bool {
			return node.Value == 100
		})
		ok(t, !success)
		assertListNodes(t, list, 1, 2, 3)
	})
}

func TestRemoveLastOccurrence(t *testing.T) {
	list := newList()

	t.Run("attempt to remove when empty", func(t *testing.T) {
		success := list.RemoveLastOccurrence(func(node *linkedlist.Node) bool {
			return node.Value == 1
		})
		ok(t, !success)
		assertListIsEmpty(t, list)
	})
	t.Run("remove tail", func(t *testing.T) {
		list := newList(1, 2, 1)
		success := list.RemoveLastOccurrence(func(node *linkedlist.Node) bool {
			return node.Value == 1
		})
		ok(t, success)
		assertListNodes(t, list, 1, 2)
	})
	t.Run("remove middle", func(t *testing.T) {
		list := newList(1, 2, 2, 3)
		success := list.RemoveLastOccurrence(func(node *linkedlist.Node) bool {
			return node.Value == 2
		})
		ok(t, success)
		assertListNodes(t, list, 1, 2, 3)
	})
	t.Run("remove head", func(t *testing.T) {
		list := newList(1, 2, 3)
		success := list.RemoveLastOccurrence(func(node *linkedlist.Node) bool {
			return node.Value == 1
		})
		ok(t, success)
		assertListNodes(t, list, 2, 3)
	})
	t.Run("remove no match", func(t *testing.T) {
		list := newList(1, 2, 3)
		success := list.RemoveLastOccurrence(func(node *linkedlist.Node) bool {
			return node.Value == 100
		})
		ok(t, !success)
		assertListNodes(t, list, 1, 2, 3)
	})
}

func TestInsertBefore(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		list := newList()
		list.InsertBefore(func(node *linkedlist.Node) bool {
			return node.Value == 100
		}, &linkedlist.Node{Value: 200})
		assertListIsEmpty(t, list)
	})
	t.Run("before head", func(t *testing.T) {
		list := newList(2)
		list.InsertBefore(func(node *linkedlist.Node) bool {
			return node.Value == 2
		}, &linkedlist.Node{Value: 1})
		assertListNodes(t, list, 1, 2)
	})
	t.Run("before middle", func(t *testing.T) {
		list := newList(1, 3, 4)
		list.InsertBefore(func(node *linkedlist.Node) bool {
			return node.Value == 3
		}, &linkedlist.Node{Value: 2})
		assertListNodes(t, list, 1, 2, 3, 4)
	})
	t.Run("before tail", func(t *testing.T) {
		list := newList(1, 2, 4)
		list.InsertBefore(func(node *linkedlist.Node) bool {
			return node.Value == 4
		}, &linkedlist.Node{Value: 3})
		assertListNodes(t, list, 1, 2, 3, 4)
	})
}

func TestInsertAfter(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		list := newList()
		list.InsertAfter(func(node *linkedlist.Node) bool {
			return node.Value == 1
		}, &linkedlist.Node{Value: 2})
		assertListIsEmpty(t, list)
	})
	t.Run("head", func(t *testing.T) {
		list := newList(1)
		list.InsertAfter(func(node *linkedlist.Node) bool {
			return node.Value == 1
		}, &linkedlist.Node{Value: 2})
		assertListNodes(t, list, 1, 2)
	})
	t.Run("middle", func(t *testing.T) {
		list := newList(1, 2, 4)
		list.InsertAfter(func(node *linkedlist.Node) bool {
			return node.Value == 2
		}, &linkedlist.Node{Value: 3})
		assertListNodes(t, list, 1, 2, 3, 4)
	})
	t.Run("tail", func(t *testing.T) {
		list := newList(1, 2, 3)
		list.InsertAfter(func(node *linkedlist.Node) bool {
			return node.Value == 3
		}, &linkedlist.Node{Value: 4})
		assertListNodes(t, list, 1, 2, 3, 4)
	})
}

func TestAddAll(t *testing.T) {
	t.Run("add to empty list", func(t *testing.T) {
		list := newList()
		list.AddAll([]*linkedlist.Node{
			&linkedlist.Node{Value: 1},
			&linkedlist.Node{Value: 2},
			&linkedlist.Node{Value: 3},
		})
		assertListNodes(t, list, 1, 2, 3)
	})
	t.Run("add to populated list", func(t *testing.T) {
		list := newList(1)
		list.AddAll([]*linkedlist.Node{
			&linkedlist.Node{Value: 1},
			&linkedlist.Node{Value: 2},
			&linkedlist.Node{Value: 3},
		})
		assertListNodes(t, list, 1, 1, 2, 3)
	})
}

func TestClear(t *testing.T) {
	list := newList(1, 2, 3)
	list.Clear()
	assertListIsEmpty(t, list)
}

func TestGet(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		list := newList()
		success, node := list.Get(0)
		ok(t, !success)
		ok(t, node == nil)
	})

	list := newList(1, 2, 3)

	t.Run("get first", func(t *testing.T) {
		success, node := list.Get(0)
		ok(t, success)
		assertNodeValue(t, 1, node)
	})

	t.Run("get second", func(t *testing.T) {
		success, node := list.Get(1)
		ok(t, success)
		assertNodeValue(t, 2, node)
	})

	t.Run("get third (last)", func(t *testing.T) {
		success, node := list.Get(2)
		ok(t, success)
		assertNodeValue(t, 3, node)
	})

	t.Run("get nil index", func(t *testing.T) {
		success, node := list.Get(3)
		ok(t, !success)
		ok(t, node == nil)
	})
}

func TestToSlice(t *testing.T) {
	list := newList(1, 2, 3)
	slice := list.ToSlice()

	collect := []int{}
	for i := 0; i < len(slice); i++ {
		collect = append(collect, slice[i].Value.(int))
	}

	assertSlicesAreEqual(t, []int{1, 2, 3}, collect)
}

func TestSet(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		list := newList()
		success := list.Set(0, &linkedlist.Node{Value: 1})
		ok(t, !success)
		assertListIsEmpty(t, list)
	})
	t.Run("one", func(t *testing.T) {
		list := newList(1)

		success := list.Set(0, &linkedlist.Node{Value: 100})
		ok(t, success)
		assertListNodes(t, list, 100)
	})
	t.Run("replace first when multiple exist in list", func(t *testing.T) {
		list := newList(1, 2)
		success := list.Set(0, &linkedlist.Node{Value: 100})
		ok(t, success)
		assertListNodes(t, list, 100, 2)
	})
	t.Run("middle", func(t *testing.T) {
		list := newList(1, 2, 3)

		success := list.Set(1, &linkedlist.Node{Value: 100})
		ok(t, success)
		assertListNodes(t, list, 1, 100, 3)
	})
	t.Run("tail", func(t *testing.T) {
		list := newList(1, 2, 3)

		success := list.Set(2, &linkedlist.Node{Value: 100})
		ok(t, success)
		assertListNodes(t, list, 1, 2, 100)
	})
}

func TestIntegration(t *testing.T) {
	list := newList(1, 2, 3)

	success, node := list.Get(0)
	assertNodeValue(t, 1, node)
	assertListNodes(t, list, 1, 2, 3)

	list.RemoveHead()
	assertListNodes(t, list, 2, 3)

	success, node = list.Get(0)
	assertNodeValue(t, 2, node)
	assertListNodes(t, list, 2, 3)

	node = list.RemoveTail()
	assertNodeValue(t, 3, node)
	assertListNodes(t, list, 2)

	list.InsertBefore(buildMatcherFn(2), &linkedlist.Node{Value: 1})
	assertListNodes(t, list, 1, 2)

	list.InsertAfter(buildMatcherFn(2), &linkedlist.Node{Value: 3})
	assertListNodes(t, list, 1, 2, 3)

	list.AddAll([]*linkedlist.Node{
		&linkedlist.Node{Value: 4},
		&linkedlist.Node{Value: 5},
		&linkedlist.Node{Value: 6},
	})
	assertListNodes(t, list, 1, 2, 3, 4, 5, 6)

	list.Add(&linkedlist.Node{Value: 7})
	assertListNodes(t, list, 1, 2, 3, 4, 5, 6, 7)

	i := list.IndexOf(func(node *linkedlist.Node) bool {
		return node.Value == 4
	})
	if i != 3 {
		t.Fatal("IndexOf failed")
	}

	success, node = list.Get(i)
	ok(t, success)
	if node.Value != 4 {
		t.Fatal("Get failed")
	}

	assertListNodes(t, list, 1, 2, 3, 4, 5, 6, 7)

	list.Add(&linkedlist.Node{Value: 1})
	assertListNodes(t, list, 1, 2, 3, 4, 5, 6, 7, 1)

	i = list.LastIndexOf(func(node *linkedlist.Node) bool {
		return node.Value == 1
	})
	ok(t, i == 7)

	success = list.RemoveLastOccurrence(func(node *linkedlist.Node) bool {
		return node.Value == 1
	})
	ok(t, success)
	assertListNodes(t, list, 1, 2, 3, 4, 5, 6, 7)

	list.AddToStart(&linkedlist.Node{Value: 1})
	assertListNodes(t, list, 1, 1, 2, 3, 4, 5, 6, 7)

	node = list.RemoveHead()
	ok(t, node.Value == 1)
	assertListNodes(t, list, 1, 2, 3, 4, 5, 6, 7)

	list.Set(3, &linkedlist.Node{Value: 1000})
	assertListNodes(t, list, 1, 2, 3, 1000, 5, 6, 7)

	list.Set(3, &linkedlist.Node{Value: 4})
	assertListNodes(t, list, 1, 2, 3, 4, 5, 6, 7)
}

func buildMatcherFn(i int) linkedlist.MatcherFn {
	return func(node *linkedlist.Node) bool {
		return node.Value == i
	}
}

func TestIndexOf(t *testing.T) {
	list := newList()

	if list.IndexOf(buildMatcherFn(1)) != -1 {
		t.Fatal("unexpected")
	}

	list.Add(&linkedlist.Node{Value: 1})
	if list.IndexOf(buildMatcherFn(1)) != 0 {
		t.Fatal("unexpected")
	}

	list.Add(&linkedlist.Node{Value: 1})
	if list.IndexOf(buildMatcherFn(1)) != 0 {
		t.Fatal("unexpected")
	}

	list.Add(&linkedlist.Node{Value: 2})
	if list.IndexOf(buildMatcherFn(2)) != 2 {
		t.Fatal("unexpected")
	}
}

func TestLastIndexOf(t *testing.T) {
	list := newList()

	if list.LastIndexOf(buildMatcherFn(1)) != -1 {
		t.Fatal("unexpected")
	}

	list.Add(&linkedlist.Node{Value: 100})
	if list.LastIndexOf(buildMatcherFn(100)) != 0 {
		t.Fatal("unexpected")
	}

	list.Add(&linkedlist.Node{Value: 100})
	if list.LastIndexOf(buildMatcherFn(100)) != 1 {
		t.Fatal("unexpected")
	}

	list.Add(&linkedlist.Node{Value: 200})
	if list.LastIndexOf(buildMatcherFn(200)) != 2 {
		t.Fatal("unexpected")
	}
}

func TestCopy(t *testing.T) {
	a := newList(1, 2, 3)
	b := a.Copy(func(node *linkedlist.Node) *linkedlist.Node {
		return &linkedlist.Node{Value: node.Value}
	})

	assertListNodes(t, a, 1, 2, 3)
	assertListNodes(t, b, 1, 2, 3)
}

// Stolen from https://github.com/stretchr/testify/blob/master/assert/assertions.go#L103
// CallerInfo returns an array of strings containing the file and line number
// of each stack frame leading from the current test to the assert call that
// failed.
func CallerInfo() []string {

	pc := uintptr(0)
	file := ""
	line := 0
	ok := false
	name := ""

	callers := []string{}
	for i := 0; ; i++ {
		pc, file, line, ok = runtime.Caller(i)
		if !ok {
			// The breaks below failed to terminate the loop, and we ran off the
			// end of the call stack.
			break
		}

		// This is a huge edge case, but it will panic if this is the case, see #180
		if file == "<autogenerated>" {
			break
		}

		f := runtime.FuncForPC(pc)
		if f == nil {
			break
		}
		name = f.Name()

		// testing.tRunner is the standard library function that calls
		// tests. Subtests are called directly by tRunner, without going through
		// the Test/Benchmark/Example function that contains the t.Run calls, so
		// with subtests we should break when we hit tRunner, without adding it
		// to the list of callers.
		if name == "testing.tRunner" {
			break
		}

		parts := strings.Split(file, "/")
		file = parts[len(parts)-1]
		if len(parts) > 1 {
			dir := parts[len(parts)-2]
			if (dir != "assert" && dir != "mock" && dir != "require") || file == "mock_test.go" {
				callers = append(callers, fmt.Sprintf("%s:%d", file, line))
			}
		}

		// Drop the package
		// segments := strings.Split(name, ".")
		// name = segments[len(segments)-1]
		// if isTest(name, "Test") ||
		// 	isTest(name, "Benchmark") ||
		// 	isTest(name, "Example") {
		// 	break
		// }
	}

	return callers
}

func newList(nums ...int) linkedlist.LinkedList {
	list := linkedlist.New()
	for _, num := range nums {
		list.Add(&linkedlist.Node{Value: num})
	}
	return list
}

func assertNodeValue(t *testing.T, expectedValue int, node *linkedlist.Node) {
	if node.Value != expectedValue {
		t.Fatal("node value is unexpected")
	}
}

func assertListIsEmpty(t *testing.T, list linkedlist.LinkedList) {
	if list.Size() != 0 {
		t.Fatal("list is not empty")
	}
	if list.Head != nil {
		t.Fatal("head is not nil")
	}
	if list.Tail != nil {
		t.Fatal("tail is not nil")
	}
}

func assertListNodes(t *testing.T, list linkedlist.LinkedList, expected ...int) {
	if list.Head == nil {
		t.Fatal("list does not have any nodes")
	}
	if list.Size() != len(expected) {
		t.Fatal("list size is unexpected")
	}
	ascending := list.Head
	descending := list.Tail
	for i := 0; i < len(expected); i++ {
		if ascending.Value != expected[i] {
			x := strings.Join(CallerInfo(), "\n\t\t\t")
			t.Fatal("ascending node value is unexpected - index: ", i, " got: ", ascending.Value, "expected: ", expected[i], x)
		}
		ascending = ascending.Next()
		if descending.Value != expected[len(expected)-1-i] {
			x := strings.Join(CallerInfo(), "\n\t\t\t")
			t.Fatal("descending node value is unexpected - index: ", i, " got: ", descending.Value, "expected: ", expected[i], x)
		}
		descending = descending.Previous()
	}
}

func ok(t *testing.T, b bool) {
	if b == false {
		t.Fatal("not ok")
	}
}

func assertSlicesAreEqual(t *testing.T, expected, got []int) {
	if len(expected) != len(got) {
		t.Fatal("slices are not of equal length")
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != got[i] {
			t.Fatal("slice values are not equal")
		}
	}
}
