package set

// SetInt maintains a unique collection of int values
type SetInt struct {
	data map[int]bool
}

// NewInt returns a new SetInt instance
func NewInt() SetInt {
	return SetInt{data: map[int]bool{}}
}

// AddInt adds the value to the collection if it does not already exist
func (s *SetInt) AddInt(v int) {
	_, ok := s.data[v]
	if !ok {
		s.data[v] = true
	}
}

// RemoveInt removes the value from the collection
func (s *SetInt) RemoveInt(v int) {
	delete(s.data, v)
}

// Size returns the total length of the collection
func (s *SetInt) Size() int {
	return len(s.data)
}

// HasInt indicates if the int exists in the collection
func (s *SetInt) HasInt(v int) bool {
	_, ok := s.data[v]
	return ok
}

// IterateInt loops over the values in the set
func (s *SetInt) IterateInt(cb func(int)) {
	for key := range s.data {
		cb(key)
	}
}

// SubsetInt tests if the specified set is a subset
func (s *SetInt) SubsetInt(b SetInt) bool {
	if s.Size() == 0 || b.Size() == 0 {
		return false
	}
	allMatch := true
	b.IterateInt(func(i int) {
		if !s.HasInt(i) {
			allMatch = false
		}
	})

	return allMatch
}

// UnionInt combines any number of sets into a new set
func UnionInt(sets ...SetInt) SetInt {
	s := NewInt()
	for i := 0; i < len(sets); i++ {
		sets[i].IterateInt(func(v int) { s.AddInt(v) })
	}
	return s
}

// IntersectionInt returns all of the values that exists in both sets as a new set
func IntersectionInt(a, b SetInt) SetInt {
	s := NewInt()
	a.IterateInt(func(v int) {
		if b.HasInt(v) {
			s.AddInt(v)
		}
	})
	return s
}
