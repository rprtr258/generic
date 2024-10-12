package fingertree

import "github.com/rprtr258/fingertree/internal"

type FingerTree[T any] struct {
	impl internal.FingerTree
}

func Foldl[T, R any](ft FingerTree[T], f func(R, any) R, initial R) R {
	return internal.Foldl(ft.impl, f, initial)
}

func Foldr[T, R any](ft FingerTree[T], f func(R, any) R, initial R) R {
	return internal.Foldr(ft.impl, f, initial)
}

func (ft FingerTree[T]) Iterl(f func(T)) {
	ft.impl.Iterl(func(a any) {
		f(a.(T))
	})
}

func (ft FingerTree[T]) Iterr(f func(T)) {
	ft.impl.Iterr(func(a any) {
		f(a.(T))
	})
}

func (ft FingerTree[T]) Pushl(d T) FingerTree[T] {
	return FingerTree[T]{ft.impl.Pushl(d)}
}

func (ft FingerTree[T]) Pushr(d T) FingerTree[T] {
	return FingerTree[T]{ft.impl.Pushr(d)}
}

func (ft FingerTree[T]) Popl() (FingerTree[T], T, bool) {
	t, x := ft.impl.Popl()
	if x == nil {
		return FingerTree[T]{t}, *new(T), false
	}

	return FingerTree[T]{t}, x.(T), true
}

func (ft FingerTree[T]) Popr() (FingerTree[T], T, bool) {
	t, x := ft.impl.Popr()
	if x == nil {
		return FingerTree[T]{t}, *new(T), false
	}

	return FingerTree[T]{t}, x.(T), true
}

func (ft FingerTree[T]) Headl() (T, bool) {
	x := ft.impl.Headl()
	if x == nil {
		return *new(T), false
	}

	return x.(T), true
}

func (ft FingerTree[T]) Headr() (T, bool) {
	x := ft.impl.Headr()
	if x == nil {
		return *new(T), false
	}

	return x.(T), true
}

func (ft FingerTree[T]) Taill() FingerTree[T] {
	return FingerTree[T]{ft.impl.Taill()}
}

func (ft FingerTree[T]) Tailr() FingerTree[T] {
	return FingerTree[T]{ft.impl.Tailr()}
}

// Concat to the right of ft
// e.g. (1).Concatl((2)) => (1, 2)
func (ft FingerTree[T]) Concat(other FingerTree[T]) FingerTree[T] {
	return FingerTree[T]{ft.impl.Concat(other.impl)}
}

func (ft FingerTree[T]) IsEmpty() bool {
	return ft.impl.IsEmpty()
}

func (ft FingerTree[T]) ToSlice() []T {
	res := []T{}
	ft.Iterl(func(t T) {
		res = append(res, t)
	})
	return res
}

func New[T any](f ...T) FingerTree[T] {
	tree := internal.Empty()
	for _, item := range f {
		tree = tree.Pushr(item)
	}
	return FingerTree[T]{tree}
}
