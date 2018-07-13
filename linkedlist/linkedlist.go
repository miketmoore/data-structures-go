package linkedlist

// LinkedList represents a linked list data structure
type LinkedList struct {
	Head *Node
	Tail *Node
	size int
}

// Node represents one link in the linked list
type Node struct {
	Value    interface{}
	next     *Node
	previous *Node
}

// Next returns the next node, if it exists
func (n *Node) Next() *Node {
	return n.next
}

// Previous returns the previous node, if it exists
func (n *Node) Previous() *Node {
	return n.previous
}

// LinkNext adds the specified node as the next node in the list
func (n *Node) LinkNext(node *Node) {
	n.next = node
}

// LinkPrevious adds the specified node as the previous node in the list
func (n *Node) LinkPrevious(node *Node) {
	n.previous = node
}

// UnlinkNext removes the node linked as "next"
func (n *Node) UnlinkNext() {
	n.next = nil
}

// UnlinkPrevious removes the node linked as "previous"
func (n *Node) UnlinkPrevious() {
	n.previous = nil
}

// New returns an empty linked list
func New() LinkedList {
	return LinkedList{}
}

// Add appends a node to the end of the list
func (l *LinkedList) Add(node *Node) {
	l.addToTail(node)
	l.size++
}

// AddToStart adds a node to the beginning of the list
func (l *LinkedList) AddToStart(node *Node) {
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	}
	head := l.Head
	if head != nil {
		tmp := l.Head
		tmp.LinkPrevious(node)
		node.LinkNext(tmp)
		l.Head = node
		l.size++
	}
}

func (l *LinkedList) iterate(currNode *Node) *Node {
	for currNode.Next() != nil {
		currNode = currNode.Next()
	}
	return currNode
}

func (l *LinkedList) addToTail(node *Node) {
	if l.Head == nil {
		l.Head = node
	} else {
		currTail := l.iterate(l.Head)
		currTail.LinkNext(node)
		node.LinkPrevious(currTail)
	}
	l.Tail = node
}

// Iterator returns an iterator instance for iterating through the list
func (l *LinkedList) Iterator() Iterator {
	return Iterator{list: l}
}

// DescendingIterator returns a descending iterator
func (l *LinkedList) DescendingIterator() Iterator {
	return Iterator{list: l, descending: true, currIndex: l.Size() - 1}
}

// Iterator represents the iterator for the list
type Iterator struct {
	currIndex  int
	currNode   *Node
	list       *LinkedList
	descending bool
}

// HasNext indicates if a node exists after the node it calls from
func (i *Iterator) HasNext() bool {
	if i.descending {
		return i.hasPrevious()
	}
	return i.hasNext()
}

// HasPrevious indicates if a node exists before the node it calls from
func (i *Iterator) HasPrevious() bool {
	if i.descending {
		return i.hasNext()
	}
	return i.hasPrevious()
}

func (i *Iterator) hasNext() bool {
	return i.list.Head != nil && i.list.Size() > i.currIndex
}

func (i *Iterator) hasPrevious() bool {
	return i.list.Head != nil && i.currIndex > -1
}

// Next returns the next node in the list
func (i *Iterator) Next() *Node {
	if i.descending {
		if i.currIndex == (i.list.Size() - 1) {
			i.currNode = i.list.Tail
			i.currIndex--
		} else if i.currNode.Previous() != nil {
			i.currNode = i.currNode.Previous()
			i.currIndex--
		}

	} else {
		if i.currIndex == 0 {
			i.currNode = i.list.Head
			i.currIndex++
		} else if i.currNode.Next() != nil {
			i.currNode = i.currNode.Next()
			i.currIndex++
		}
	}
	return i.currNode
}

// RemoveHead removes the first node from the list
func (l *LinkedList) RemoveHead() *Node {
	if l.Head != nil {
		removed := l.Head
		if l.Head.Next() != nil {
			l.Head = l.Head.Next()
		} else {
			l.Head = nil
		}
		l.size--
		if l.size == 0 {
			l.Head = nil
			l.Tail = nil
		}
		return removed
	}
	return nil
}

// RemoveTail removes the last node from the list
func (l *LinkedList) RemoveTail() *Node {
	if l.Head == nil {
		return nil
	}
	var prev *Node
	node := l.Head
	for node.Next() != nil {
		prev = node
		node = node.Next()
	}
	var removed *Node
	if l.Tail != nil {
		removed = l.Tail
	}
	if prev != nil {
		prev.UnlinkNext()
		l.Tail = prev
		l.size--
	} else {
		l.Head = nil
		l.Tail = nil
		l.size--
	}
	return removed
}

// MatcherFn represents a matcher
type MatcherFn = func(*Node) bool

// RemoveFirstOccurrence removes the first occurence of the value in the list
func (l *LinkedList) RemoveFirstOccurrence(matcher MatcherFn) bool {
	// if the value removed is the last.. it won't have next
	if l.Head == nil {
		return false
	}
	var prev *Node
	node := l.Head
	for node.Next() != nil {
		if matcher(node) {
			if prev == nil {
				// replace head with next node
				l.Head = node.Next()
			} else {
				// replace prev-next with next
				prev.LinkNext(node.Next())
				node.Next().LinkPrevious(prev)
			}
			l.size--
			return true
		}
		prev = node
		node = node.Next()
	}
	if matcher(node) {
		prev.UnlinkNext()
		l.size--
		l.Tail = prev
		return true
	}
	return false
}

// RemoveLastOccurrence removes the last occurence of the node in the list
func (l *LinkedList) RemoveLastOccurrence(matcher MatcherFn) bool {
	if l.Head == nil {
		return false
	}
	node := l.Tail
	var prev *Node
	i := l.Size() - 1
	for node.Previous() != nil {
		if matcher(node) {
			if i == (l.Size() - 1) {
				// tail
				if node.Previous() != nil {
					node.Previous().UnlinkNext()
				}
				l.Tail = node.Previous()
			} else if i > 0 {
				// middle
				node.Previous().LinkNext(prev)
				node.Next().LinkPrevious(node.Previous())
			}
			l.size--
			return true
		}
		prev = node
		node = node.Previous()
		i--
	}
	if matcher(node) {
		// head
		l.Head = node.Next()
		l.size--
		return true
	}
	return false
}

// Find finds the first occurence of the value and returns the node
func (l *LinkedList) Find(matcher func(node *Node) bool) (bool, *Node) {
	if l.Head == nil {
		return false, nil
	}
	node := l.Head
	for node.Next() != nil {
		if matcher(node) {
			return true, node
		}
		node = node.Next()
	}
	if matcher(node) {
		return true, node
	}
	return false, nil
}

// InsertBefore inserts a new node before the matched node
func (l *LinkedList) InsertBefore(matcher MatcherFn, new *Node) {
	if l.Head == nil {
		return
	}
	if matcher(l.Head) {
		l.AddToStart(new)
		return
	}
	var prev *Node
	node := l.Head
	for node.Next() != nil {
		if matcher(node) {
			if prev != nil {
				prev.LinkNext(new)
				new.LinkPrevious(prev)
				new.LinkNext(node)
				node.LinkPrevious(new)
			}
			l.size++
			return
		}
		prev = node
		node = node.Next()
	}
	if matcher(node) {
		if prev != nil {
			prev.LinkNext(new)
			new.LinkPrevious(prev)
			new.LinkNext(node)
			node.LinkPrevious(new)
		}
		l.size++
		return
	}
}

// InsertAfter inserts a new node before the matched node
func (l *LinkedList) InsertAfter(matcher MatcherFn, new *Node) {
	if l.Head == nil {
		return
	}
	if matcher(l.Head) {
		if l.Head.Next() != nil {
			// head matches and is already linked to next node
			// cache current next node
			tmp := l.Head.Next()
			// link new node as next
			l.Head.LinkNext(new)
			new.LinkPrevious(l.Head)

			new.LinkNext(tmp)
			tmp.LinkPrevious(new)
		} else {
			// head matches and is not linked to a next node
			l.Head.LinkNext(new)
			new.LinkPrevious(l.Head)
			l.Tail = new
		}
		l.size++
		return
	}
	node := l.Head.Next()
	for node.Next() != nil {
		if matcher(node) {
			tmp := node.Next()
			node.LinkNext(new)
			new.LinkPrevious(node)

			node.Next().LinkNext(tmp)
			tmp.LinkPrevious(node.Next())
			l.size++
			return
		}
		node = node.Next()
	}
	if matcher(node) {
		node.LinkNext(new)
		new.LinkPrevious(node)
		l.Tail = node.Next()
		l.size++
		return
	}
}

// AddAll appends items to the end of the list
func (l *LinkedList) AddAll(all []*Node) {
	for i := 0; i < len(all); i++ {
		l.Add(all[i])
	}
}

// Clear removes all items from the list
func (l *LinkedList) Clear() {
	l.Head = nil
	l.Tail = nil
	l.size = 0
}

// Get returns the node at the specified index
func (l *LinkedList) Get(index int) (bool, *Node) {
	if l.Head == nil {
		return false, nil
	}
	i := 0
	node := l.Head
	for node.Next() != nil {
		if i == index {
			return true, node
		}
		node = node.Next()
		i++
	}
	if i == index {
		return true, node
	}
	return false, nil
}

// Size returns the total number of nodes in the list
func (l *LinkedList) Size() int {
	return l.size
}

// ToSlice returns a slice of the nodes in this list
func (l *LinkedList) ToSlice() []*Node {
	slice := make([]*Node, l.Size())
	it := l.Iterator()
	i := 0
	for it.HasNext() {
		slice[i] = it.Next()
		i++
	}
	return slice
}

// Set replaces the node at the specified index
func (l *LinkedList) Set(index int, new *Node) bool {
	if l.Head == nil {
		return false
	}
	if index == 0 {
		if l.Head == l.Tail {
			l.Head = new
			l.Tail = new
			return true
		}
		new.LinkNext(l.Head.Next())
		l.Head.Next().LinkPrevious(new)
		l.Head = new
		return true
	}
	node := l.Head
	i := 0
	for node.Next() != nil {
		if index == i {
			if node.Previous() != nil {
				new.LinkPrevious(node.Previous())
				node.Previous().LinkNext(new)
			}
			new.LinkNext(node.Next())
			node.Next().LinkPrevious(new)
			return true
		}
		i++
		node = node.Next()
	}
	if index == i {
		if node.Previous() != nil {
			new.LinkPrevious(node.Previous())
			node.Previous().LinkNext(new)
		}
		l.Tail = new
		return true
	}
	return false
}

// IndexOf returns the first index of the node in the list
func (l *LinkedList) IndexOf(matcher func(node *Node) bool) int {
	if l.Head == nil {
		return -1
	}
	i := 0
	node := l.Head
	for node.Next() != nil {
		if matcher(node) {
			return i
		}
		i++
		node = node.Next()
	}
	if matcher(node) {
		return i
	}
	return -1
}

// LastIndexOf returns the last index of the value in the list
func (l *LinkedList) LastIndexOf(matcher func(node *Node) bool) int {
	if l.Head == nil {
		return -1
	}
	lastMatchIndex := -1
	i := 0
	node := l.Head
	for node.Next() != nil {
		if matcher(node) {
			lastMatchIndex = i
		}
		i++
		node = node.Next()
	}
	if matcher(node) {
		lastMatchIndex = i
	}
	if lastMatchIndex > -1 {
		return lastMatchIndex
	}
	return -1
}

// Copy returns a copy of this linked list
func (l *LinkedList) Copy(copy func(node *Node) *Node) LinkedList {
	new := New()
	it := l.Iterator()
	for it.HasNext() {
		node := it.Next()
		new.Add(copy(node))
	}
	return new
}
