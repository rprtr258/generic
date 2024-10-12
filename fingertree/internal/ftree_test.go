package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func toSlice(t FingerTree) []any {
	return Foldl(t, func(a []any, b any) []any {
		return append(a, b)
	}, make([]any, 0))
}

func concat(ts ...FingerTree) []any {
	return _foldl(ts, func(a []any, b FingerTree) []any {
		return append(a, toSlice(b)...)
	}, []any{})
}

func TestEmptyFoldl(t *testing.T) {
	n := empty()
	add := func(a []any, b any) []any {
		return append(a, b)
	}
	r := Foldl(n, add, []any{})
	assert.Equal(t, []any{}, r)
}

func TestEmptyFoldr(t *testing.T) {
	n := empty()
	add := func(a []any, b any) []any {
		return append(a, b)
	}
	r := Foldr(n, add, []any{})
	assert.Equal(t, []any{}, r)
}

func TestEmptyIterr(t *testing.T) {
	n := empty()
	sum := 0
	add := func(b any) {
		sum += b.(int)
	}
	n.Iterr(add)
	assert.Equal(t, 0, sum)
}

func TestEmptyIterl(t *testing.T) {
	n := empty()
	sum := 0
	add := func(b any) {
		sum += b.(int)
	}
	n.Iterl(add)
	assert.Equal(t, 0, sum)
}

func TestEmptyPushl(t *testing.T) {
	v := (empty()).Pushl(1)
	assert.Equal(t, []any{1}, toSlice(v))
}

func TestEmptyPopl(t *testing.T) {
	n := empty()
	r, e := n.Popl()

	assert.True(t, r.isA)
	assert.Nil(t, e)
}

func TestEmptyPopr(t *testing.T) {
	n := empty()
	r, e := n.Popr()

	assert.True(t, r.isA)
	assert.Nil(t, e)
}

func TestEmptyPushr(t *testing.T) {
	v := (empty()).Pushr(1)
	assert.Equal(t, []any{1}, toSlice(v))
}

func TestEmptyHeadr(t *testing.T) {
	v := (empty()).Headr()
	assert.Nil(t, v)
}

func TestEmptyTailr(t *testing.T) {
	v := (empty()).Tailr()
	assert.True(t, v.isA)
}

func TestEmptyHeadl(t *testing.T) {
	v := (empty()).Headl()
	assert.Nil(t, v)
}

func TestEmptyTaill(t *testing.T) {
	v := (empty()).Taill()
	assert.True(t, v.isA)
}

func TestEmptyIsEmpty(t *testing.T) {
	v := empty()
	assert.True(t, v.IsEmpty())
}

func TestEmptyConcatl(t *testing.T) {
	e := empty()
	s := e.Pushl(1)
	kal := s.Pushl(2)

	o := empty()

	assert.Equal(t, toSlice(kal), toSlice(kal.Concat(e)))
	assert.Equal(t, toSlice(s), toSlice(s.Concat(e)))
	assert.Equal(t, toSlice(o), toSlice(o.Concat(e)))
}

func TestEmptyConcat(t *testing.T) {
	e := empty()
	s := e.Pushl(1)
	kal := s.Pushl(2)

	o := e

	assert.Equal(t, toSlice(kal), toSlice(e.Concat(kal)))
	assert.Equal(t, toSlice(s), toSlice(e.Concat(s)))
	assert.Equal(t, toSlice(o), toSlice(e.Concat(o)))
}

func TestSinglePushl(t *testing.T) {
	r := single(1).Pushl(2)
	assert.Equal(t, []any{2, 1}, toSlice(r))
}

func TestSinglePopl(t *testing.T) {
	n := single(1)
	r, e := n.Popl()
	assert.True(t, r.isA)
	assert.Equal(t, 1, e)
}

func TestSinglePushr(t *testing.T) {
	n := single(1)
	r := n.Pushr(2)
	assert.Equal(t, []any{1, 2}, toSlice(r))
}

func TestSingleFoldl(t *testing.T) {
	n := single(1)
	r := Foldl(n, func(a []any, b any) []any {
		return append(a, b)
	}, []any{})
	assert.Equal(t, r, []any{1})
}

func TestSingleFoldr(t *testing.T) {
	n := single(1)
	r := Foldr(n, func(a []any, b any) []any {
		return append(a, b)
	}, []any{})
	assert.Equal(t, r, []any{1})
}

func TestSingleIterr(t *testing.T) {
	n := single(1)
	sum := 0
	add := func(b any) {
		sum += b.(int)
	}
	n.Iterr(add)
	assert.Equal(t, 1, sum)
}

func TestSingleIterl(t *testing.T) {
	n := single(1)
	sum := 0
	add := func(b any) {
		sum += b.(int)
	}
	n.Iterl(add)
	assert.Equal(t, 1, sum)
}

func TestSingleHeadr(t *testing.T) {
	v := single(1).Headr()
	if v != 1 {
		t.Error(fmt.Sprintf("single{1}.Headr() should be 1, got %v", v))
	}
}

func TestSingleTailr(t *testing.T) {
	v := single(1).Tailr()
	if !v.IsEmpty() {
		t.Error(fmt.Sprintf("single{1}.Tailr() should be empty, got %v", v))
	}
}

func TestSingleHeadl(t *testing.T) {
	v := single(1).Headl()
	if v != 1 {
		t.Error(fmt.Sprintf("single{1}.Headl() should be 1, got %v", v))
	}
}

func TestSingleTaill(t *testing.T) {
	v := single(1).Taill()
	if !v.IsEmpty() {
		t.Error(fmt.Sprintf("single{1}.Taill() should be empty, got %v", v))
	}
}

func TestSingleIsEmpty(t *testing.T) {
	v := single(1)
	if v.IsEmpty() {
		t.Error("Expected &single{1}.IsEmpty() to be false")
	}
}

func TestSingleConcat(t *testing.T) {
	s := empty()
	s1 := s.Pushl(1)
	s21 := s1.Pushl(2)
	s3 := empty().Pushl(3)

	assert.Equal(t, concat(s21, s1), toSlice(s21.Concat(s1)))
	assert.Equal(t, concat(s21, s3), toSlice(s21.Concat(s3)))
	assert.Equal(t, concat(s3, s1), toSlice(s3.Concat(s1)))
	assert.Equal(t, concat(s, s1), toSlice(s.Concat(s1)))
	assert.Equal(t, concat(s1, s21), toSlice(s1.Concat(s21)))
	assert.Equal(t, concat(s3, s21), toSlice(s3.Concat(s21)))
	assert.Equal(t, concat(s1, s3), toSlice(s1.Concat(s3)))
	assert.Equal(t, concat(s1, s), toSlice(s1.Concat(s)))
}

func TestFTreeFoldl(t *testing.T) {
	var n FingerTree = empty()
	for i := 0; i < 20; i++ {
		n = n.Pushr(i)
	}

	r := Foldl(n, func(a []any, b any) []any {
		return append(a, b)
	}, []any{})
	assert.Equal(t, []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, r)
}

func TestFTreeFoldr(t *testing.T) {
	var n FingerTree = empty()
	for i := 0; i < 20; i++ {
		n = n.Pushl(i)
	}

	r := Foldr(n, func(a []any, b any) []any {
		return append(a, b)
	}, []any{})
	expect := []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	assert.Equal(t, r, expect)
}

func TestFTreePushl(t *testing.T) {
	var n FingerTree = single(0)
	for i := 1; i < 20; i++ {
		n = n.Pushl(i)
	}

	assert.Equal(t, toSlice(n), []any{19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
}

func TestFTreePopl(t *testing.T) {
	var n FingerTree = empty()
	for i := 0; i < 20; i++ {
		n = n.Pushl(i)
	}

	var e any
	for i := 19; i >= 0; i-- {
		n, e = n.Popl()
		if e != i {
			t.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}

	for i := 0; i < 22; i++ {
		n = n.Pushr(i)
	}
	for i := 0; i < 22; i++ {
		n, e = n.Popl()
		if e != i {
			t.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}
}

func TestFTreePopr(t *testing.T) {
	var n FingerTree = empty()
	for i := 0; i < 20; i++ {
		n = n.Pushr(i)
	}

	var e any
	for i := 19; i >= 0; i-- {
		n, e = n.Popr()
		if e != i {
			t.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}

	for i := 0; i < 22; i++ {
		n = n.Pushl(i)
	}
	for i := 0; i < 22; i++ {
		n, e = n.Popr()
		if e != i {
			t.Error(fmt.Sprintf("Expected pop result to be %v, got %v", i, e))
		}
	}
}

func TestFTreePushr(t *testing.T) {
	n := FingerTree(single(0))
	for i := 1; i < 20; i++ {
		n = n.Pushr(i)
	}

	assert.Equal(t, []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, toSlice(n))
}

func TestFTreeIterl(t *testing.T) {
	var n FingerTree = empty()
	for i := 0; i < 10; i++ {
		n = n.Pushl(i)
	}

	sum := 0
	n.Iterl(func(d any) {
		sum += d.(int)
	})
	if sum != 45 {
		t.Error(fmt.Sprintf("Expected n.Iterl to result in sum 110, got %v", sum))
	}
}

func TestFTreeIterr(t *testing.T) {
	var n FingerTree = empty()
	for i := 0; i < 10; i++ {
		n = n.Pushl(i)
	}

	sum := 0
	n.Iterr(func(d any) {
		sum += d.(int)
	})
	if sum != 45 {
		t.Error(fmt.Sprintf("Expected n.Iterl to result in sum 110, got %v", sum))
	}
}

func TestFTreeHeadr(t *testing.T) {
	v := (empty()).Pushr(1).Pushr(2)
	r := v.Headr()
	if r != 2 {
		t.Error(fmt.Sprintf("ftree{1 2}.Headr() should be 2, got %v", r))
	}

	v = (empty()).Pushl(1).Pushl(2)
	r = v.Headr()
	if r != 1 {
		t.Error(fmt.Sprintf("ftree{1 2}.Headr() should be 2, got %v", r))
	}
}

func TestFTreeHeadl(t *testing.T) {
	v := triple([]any{1}, []any{2}, empty()).Headl()
	if v != 1 {
		t.Error(fmt.Sprintf("ftree{1 2}.Headr() should be 1, got %v", v))
	}
}

func TestFTreeConcatl_case3(t *testing.T) {
	e := empty()
	s := e.Pushl(1)

	kal := FingerTree(empty())
	for i := 0; i < 25; i++ {
		kal = kal.Pushl(i)
	}
	o := FingerTree(empty())
	for i := 0; i < 105; i++ {
		o = o.Pushl(i)
	}
	assert.Equal(t, append(toSlice(o), toSlice(kal)...), toSlice(o.Concat(kal)))
	assert.Equal(t, append(toSlice(s), toSlice(kal)...), toSlice(s.Concat(kal)))
	assert.Equal(t, append(toSlice(e), toSlice(kal)...), toSlice(e.Concat(kal)))
}

func TestFTreeConcatl_case4(t *testing.T) {
	s := empty()
	s1 := s.Pushl(1)

	s012 := FingerTree(empty())
	for i := 0; i < 3; i++ {
		s012 = s012.Pushl(i)
	}
	s210 := FingerTree(empty())
	for i := 0; i < 3; i++ {
		s210 = s210.Pushl(i)
	}
	assert.Equal(t, append(toSlice(s210), toSlice(s012)...), toSlice(s210.Concat(s012)))
	assert.Equal(t, append(toSlice(s1), toSlice(s012)...), toSlice(s1.Concat(s012)))
	assert.Equal(t, append(toSlice(s), toSlice(s012)...), toSlice(s.Concat(s012)))
}

func TestFTreeConcatl(t *testing.T) {
	e := empty()
	s := e.Pushl(1)
	var kal FingerTree = empty()
	for i := 0; i < 25; i++ {
		kal = kal.Pushl(i)
	}

	var o FingerTree = empty()
	for i := 0; i < 5; i++ {
		o = o.Pushl(i)
	}

	testCombinations := func() {
		assert.Equal(t, append(toSlice(o), toSlice(kal)...), toSlice(o.Concat(kal)))
		assert.Equal(t, append(toSlice(s), toSlice(kal)...), toSlice(s.Concat(kal)))
		assert.Equal(t, append(toSlice(e), toSlice(kal)...), toSlice(e.Concat(kal)))
	}

	testCombinations()

	kal = empty()
	o = empty()
	for i := 0; i < 5; i++ {
		kal = kal.Pushl(i)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushl(i)
	}
	testCombinations()

	kal = empty()
	o = empty()
	for i := 0; i < 105; i++ {
		kal = kal.Pushl(i)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushl(i)
	}
	testCombinations()
}

func TestFTreeConcat(t *testing.T) {
	e := empty()
	s := e.Pushr(1)
	kal := FingerTree(empty())
	for i := 0; i < 25; i++ {
		kal = kal.Pushr(i)
	}

	var o FingerTree = empty()
	for i := 0; i < 5; i++ {
		o = o.Pushr(i)
	}

	testCombinations := func() {
		expected := append(toSlice(kal), toSlice(o)...)
		r := kal.Concat(o)
		assert.Equal(t, expected, toSlice(r))

		expected = append(toSlice(kal), toSlice(s)...)
		r = kal.Concat(s)
		assert.Equal(t, expected, toSlice(r))

		expected = append(toSlice(kal), toSlice(e)...)
		r = kal.Concat(e)
		assert.Equal(t, expected, toSlice(r))
	}

	testCombinations()

	kal = empty()
	o = empty()
	for i := 0; i < 5; i++ {
		kal = kal.Pushr(i)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushr(i)
	}
	testCombinations()

	kal = empty()
	o = empty()
	for i := 0; i < 25; i++ {
		kal = kal.Pushr(i)
	}
	for i := 0; i < 105; i++ {
		o = o.Pushr(i)
	}
	testCombinations()

	kal = empty()
	o = empty()
	for i := 0; i < 105; i++ {
		kal = kal.Pushr(i)
	}
	for i := 0; i < 25; i++ {
		o = o.Pushr(i)
	}
	testCombinations()
}
