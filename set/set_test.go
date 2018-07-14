package set_test

import (
	"testing"

	"github.com/miketmoore/data-structures-go/set"
)

func TestIntegration(t *testing.T) {
	s := set.NewInt()

	s.AddInt(1)
	ok(t, s.Size() == 1)
	s.AddInt(1)
	ok(t, s.Size() == 1)

	s.AddInt(2)
	ok(t, s.Size() == 2)
	s.AddInt(2)
	ok(t, s.Size() == 2)

	s.AddInt(3)
	ok(t, s.Size() == 3)
	s.AddInt(3)
	ok(t, s.Size() == 3)

	ok(t, s.HasInt(1))
	ok(t, s.HasInt(2))
	ok(t, s.HasInt(3))
	ok(t, !s.HasInt(0))
	ok(t, !s.HasInt(100))

	s.RemoveInt(1)
	ok(t, !s.HasInt(1))
	ok(t, s.HasInt(2))
	ok(t, s.HasInt(3))
	ok(t, s.Size() == 2)

	s.RemoveInt(2)
	ok(t, !s.HasInt(1))
	ok(t, !s.HasInt(2))
	ok(t, s.HasInt(3))
	ok(t, s.Size() == 1)

	s.RemoveInt(3)
	ok(t, !s.HasInt(1))
	ok(t, !s.HasInt(2))
	ok(t, !s.HasInt(3))
	ok(t, s.Size() == 0)
}

func TestUnion(t *testing.T) {
	a := set.NewInt()
	a.AddInt(1)
	a.AddInt(2)

	b := set.NewInt()
	b.AddInt(1)
	b.AddInt(3)

	c := set.NewInt()
	c.AddInt(3)
	c.AddInt(4)

	d := set.UnionInt(a, b, c)

	ok(t, d.Size() == 4)
	ok(t, d.HasInt(1))
	ok(t, d.HasInt(2))
	ok(t, d.HasInt(3))
	ok(t, d.HasInt(4))
}

func TestIntersection(t *testing.T) {
	a := set.NewInt()
	a.AddInt(1)
	a.AddInt(2)
	a.AddInt(3)

	b := set.NewInt()
	b.AddInt(2)
	b.AddInt(3)
	b.AddInt(4)

	c := set.IntersectionInt(a, b)

	ok(t, c.Size() == 2)
	ok(t, !c.HasInt(1))
	ok(t, c.HasInt(2))
	ok(t, c.HasInt(3))
	ok(t, !c.HasInt(4))
}

func TestSubset(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		a := set.NewInt()
		a.AddInt(1)
		a.AddInt(2)
		a.AddInt(3)

		b := set.NewInt()
		b.AddInt(2)
		b.AddInt(3)

		isSubset := a.SubsetInt(b)
		ok(t, isSubset)
	})
	t.Run("failure", func(t *testing.T) {
		a := set.NewInt()
		a.AddInt(1)
		a.AddInt(2)
		a.AddInt(3)

		b := set.NewInt()
		a.AddInt(1)
		b.AddInt(2)
		b.AddInt(3)
		b.AddInt(4)

		isSubset := a.SubsetInt(b)
		ok(t, !isSubset)
	})
	t.Run("failure - both empty", func(t *testing.T) {
		a := set.NewInt()
		b := set.NewInt()

		isSubset := a.SubsetInt(b)
		ok(t, !isSubset)
	})
	t.Run("failure - a empty", func(t *testing.T) {
		a := set.NewInt()
		b := set.NewInt()
		b.AddInt(1)

		isSubset := a.SubsetInt(b)
		ok(t, !isSubset)
	})
	t.Run("failure - b empty", func(t *testing.T) {
		a := set.NewInt()
		a.AddInt(1)
		b := set.NewInt()

		isSubset := a.SubsetInt(b)
		ok(t, !isSubset)
	})
}

func TestDifferenceInt(t *testing.T) {
	a := set.NewInt()
	a.AddInt(1)
	a.AddInt(2)
	a.AddInt(3)

	b := set.NewInt()
	b.AddInt(2)
	b.AddInt(3)
	b.AddInt(4)

	c := set.DifferenceInt(a, b)
	ok(t, c.Size() == 2)
	ok(t, c.HasInt(1))
	ok(t, !c.HasInt(2))
	ok(t, !c.HasInt(3))
	ok(t, c.HasInt(4))
}

func ok(t *testing.T, v bool) {
	if v == false {
		t.Fatal("not ok")
	}
}
