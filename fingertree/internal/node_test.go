package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode2Foldl(t *testing.T) {
	n := node2(1, 2)
	add := func(a, b any) any {
		return any(a.(int) + b.(int))
	}
	r := nodeFoldl(n, add, 0)
	if r != 3 {
		t.Errorf("Expected n.Foldl(func (a, b uint) { return a + b }, 0) to return 3, got %d", r.(int))
	}
}

func TestNode3Foldl(t *testing.T) {
	n := node3(1, 2, 3)
	add := func(a, b any) any {
		return any(a.(int) + b.(int))
	}
	r := nodeFoldl(n, add, 0)
	if r != 6 {
		t.Errorf("Expected n.Foldl(func (a, b uint) { return a + b }, 0) to return 3, got %d", r.(uint))
	}
}

func TestNode2Foldr(t *testing.T) {
	n := node2(1, 2)
	r := nodeFoldr(n, func(a, b any) any {
		return any(a.(int) + b.(int))
	}, 0)
	if r != 3 {
		t.Errorf("Expected n.Foldr(func (a, b uint) { return a + b }, 0) to return 3, got %d", r.(int))
	}
}

func TestNode3Foldr(t *testing.T) {
	n := node3(1, 2, 3)
	r := nodeFoldr(n, func(a, b any) any {
		return any(a.(int) + b.(int))
	}, 0)
	assert.Equal(t, 6, r)
}

func TestNode2Iterr(t *testing.T) {
	n := node2(1, 2)
	sum := 0
	add := func(b any) {
		sum += b.(int)
	}
	n.Iterr(add)
	if sum != 3 {
		t.Errorf("Expected n.Iterr(func (a, b uint) { return a + b }, 0) to return 3, got %d", sum)
	}
}

func TestNode3Iterr(t *testing.T) {
	n := node3(1, 2, 3)
	sum := 0
	add := func(b any) {
		sum += b.(int)
	}
	n.Iterr(add)
	if sum != 6 {
		t.Errorf("Expected n.Iterr(func (a, b uint) { return a + b }, 0) to return 3, got %d", sum)
	}
}

func TestNode2Iterl(t *testing.T) {
	n := node2(1, 2)
	sum := 0
	add := func(b any) {
		sum += b.(int)
	}
	n.Iterl(add)
	if sum != 3 {
		t.Errorf("Expected n.Iterl(func (a, b uint) { return a + b }, 0) to return 3, got %d", sum)
	}
}

func TestNode3Iterl(t *testing.T) {
	n := node3(1, 2, 3)
	sum := 0
	add := func(b any) {
		sum += b.(int)
	}
	n.Iterl(add)
	if sum != 6 {
		t.Errorf("Expected n.Iterl(func (a, b uint) { return a + b }, 0) to return 3, got %d", sum)
	}
}
