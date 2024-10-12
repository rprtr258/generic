package fingertree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFTreeTailr(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := New(xs...)
	b := a.Tailr()
	c := b.ToSlice()
	assert.Equal(t, xs[:len(xs)-1], c)
}

func TestFTreeTaill(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, xs[1:], New(xs...).Taill().ToSlice())
}

func TestFTreeIsEmpty(t *testing.T) {
	v := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	assert.False(t, v.IsEmpty())
}
