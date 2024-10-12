package internal

import "unsafe"

type union[A, B any] struct {
	data []byte // of constant len=max(sizeof(A), sizeof(B))
}

func newUnionA[A, B any](a A) union[A, B] {
	u := union[A, B]{make([]byte, max(unsafe.Sizeof(*new(A)), unsafe.Sizeof(*new(B))))}
	*u.A() = a
	return u
}

func newUnionB[A, B any](b B) union[A, B] {
	u := union[A, B]{make([]byte, max(unsafe.Sizeof(*new(A)), unsafe.Sizeof(*new(B))))}
	*u.B() = b
	return u
}

func (u union[A, B]) A() *A {
	return (*A)(unsafe.Pointer(unsafe.SliceData(u.data)))
}

func (u union[A, B]) B() *B {
	return (*B)(unsafe.Pointer(unsafe.SliceData(u.data)))
}

type union3[A, B, C any] struct {
	data []byte // of constant len=max(sizeof(A), sizeof(B), sizeof(C))
}

func newUnion3A[A, B, C any](a A) union3[A, B, C] {
	u := union3[A, B, C]{make([]byte, max(unsafe.Sizeof(*new(A)), unsafe.Sizeof(*new(B)), unsafe.Sizeof(*new(C))))}
	*u.A() = a
	return u
}

func newUnion3B[A, B, C any](b B) union3[A, B, C] {
	u := union3[A, B, C]{make([]byte, max(unsafe.Sizeof(*new(A)), unsafe.Sizeof(*new(B)), unsafe.Sizeof(*new(C))))}
	*u.B() = b
	return u
}

func newUnion3C[A, B, C any](c C) union3[A, B, C] {
	u := union3[A, B, C]{make([]byte, max(unsafe.Sizeof(*new(A)), unsafe.Sizeof(*new(B)), unsafe.Sizeof(*new(C))))}
	*u.C() = c
	return u
}

func (u union3[A, B, C]) A() *A {
	return (*A)(unsafe.Pointer(unsafe.SliceData(u.data)))
}

func (u union3[A, B, C]) B() *B {
	return (*B)(unsafe.Pointer(unsafe.SliceData(u.data)))
}

func (u union3[A, B, C]) C() *C {
	return (*C)(unsafe.Pointer(unsafe.SliceData(u.data)))
}

type unionTag[A, B any] struct {
	union[A, B]
	isB bool
}

func NewUnionTagA[A, B any](a A) unionTag[A, B] {
	return unionTag[A, B]{
		union: newUnionA[A, B](a),
		isB:   false,
	}
}

func NewUnionTagB[A, B any](b B) unionTag[A, B] {
	return unionTag[A, B]{
		union: newUnionB[A, B](b),
		isB:   true,
	}
}

func (u unionTag[A, B]) A() *A {
	return u.union.A()
}

func (u unionTag[A, B]) B() *B {
	return u.union.B()
}

func (u unionTag[A, B]) Switch(
	fa func(*A),
	fb func(*B),
) {
	if u.isB {
		fb(u.B())
	} else {
		fa(u.A())
	}
}

func Switch[A, B, R any](
	u unionTag[A, B],
	fa func(A) R,
	fb func(B) R,
) R {
	var res R
	switch {
	case u.isB:
		res = fb(*u.B())
	default:
		res = fa(*u.A())
	}
	return res
}

type unionTag3[A, B, C any] struct {
	union3[A, B, C]
	isA, isB bool
}

func newUnionTag3A[A, B, C any](a A) unionTag3[A, B, C] {
	return unionTag3[A, B, C]{
		union3: newUnion3A[A, B, C](a),
		isA:    true,
		isB:    false,
	}
}

func newUnionTag3B[A, B, C any](b B) unionTag3[A, B, C] {
	return unionTag3[A, B, C]{
		union3: newUnion3B[A, B, C](b),
		isA:    false,
		isB:    true,
	}
}

func newUnionTag3C[A, B, C any](c C) unionTag3[A, B, C] {
	return unionTag3[A, B, C]{
		union3: newUnion3C[A, B, C](c),
		isA:    false,
		isB:    false,
	}
}

func (u unionTag3[A, B, C]) A() *A {
	return u.union3.A()
}

func (u unionTag3[A, B, C]) B() *B {
	return u.union3.B()
}

func (u unionTag3[A, B, C]) C() *C {
	return u.union3.C()
}

func (u unionTag3[A, B, C]) Switch(
	fa func(*A),
	fb func(*B),
	fc func(*C),
) {
	switch {
	case u.isA:
		fa(u.A())
	case u.isB:
		fb(u.B())
	default:
		fc(u.C())
	}
}

func Switch3[A, B, C, R any](
	u unionTag3[A, B, C],
	fa func(A) R,
	fb func(B) R,
	fc func(C) R,
) R {
	var res R
	switch {
	case u.isA:
		res = fa(*u.A())
	case u.isB:
		res = fb(*u.B())
	default:
		res = fc(*u.C())
	}
	return res
}
