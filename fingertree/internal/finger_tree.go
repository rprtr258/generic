package internal

type FingerTree struct {
	unionTag3[struct{}, any, *ftree]
}

func _foldl[T, R any](s []T, f func(R, T) R, init R) R {
	v := init
	for _, x := range s {
		v = f(v, x)
	}
	return v
}

func Foldl[R any](ft FingerTree, f func(R, any) R, initial R) R {
	return Switch3(
		ft.unionTag3,
		func(struct{}) R { return initial },
		func(a any) R { return f(initial, a) },
		func(t *ftree) R {
			a := _foldl(t.left, f, initial)
			b := Foldl(t.child, func(init R, data any) R {
				return nodeFoldl(*data.(*node), f, init)
			}, a)
			return _foldl(t.right, f, b)
		},
	)
}

func _foldr[R any](s []any, f func(R, any) R, init R) R {
	v := init
	for i := range s {
		v = f(v, s[len(s)-1-i])
	}
	return v
}

func Foldr[R any](ft FingerTree, f func(R, any) R, initial R) R {
	return Switch3(
		ft.unionTag3,
		func(struct{}) R { return initial },
		func(a any) R { return f(initial, a) },
		func(t *ftree) R {
			a := _foldr(t.right, f, initial)
			b := Foldr(t.child, func(init R, data any) R {
				return nodeFoldr(*data.(*node), f, init)
			}, a)
			return _foldr(t.left, f, b)
		},
	)
}

func Empty() FingerTree {
	return empty()
}

// Transform a slice of elements into a slice of nodes
func nodes(xs []any) []any {
	if len(xs) == 1 {
		panic("Can't make a node from a single element.")
	}
	if len(xs) == 2 {
		n := node2(xs[0], xs[1])
		return []any{&n}
	}
	if len(xs) == 3 {
		n := node3(xs[0], xs[1], xs[2])
		return []any{&n}
	}
	if len(xs) == 4 {
		n1, n2 := node2(xs[0], xs[1]), node2(xs[2], xs[3])
		return []any{&n1, &n2}
	}
	if len(xs) > 4 {
		return append(nodes(xs[:3]), nodes(xs[3:])...)
	}
	return []any{}
}

// Join two finger trees with a 'glue' slice between them
// Normally calling Concat or Concatl will be more useful
func glue(l FingerTree, c []any, r FingerTree) FingerTree {
	pushl := func(a FingerTree, s []any) FingerTree {
		m := a
		for i := range s {
			m = m.Pushl(s[len(s)-1-i])
		}
		return m
	}

	pushr := func(a FingerTree, s []any) FingerTree {
		m := a
		for _, t := range s {
			m = m.Pushr(t)
		}
		return m
	}

	switch {
	// If either branch is empty, it can be ignored
	case l.IsEmpty():
		return pushl(r, c)
	case r.IsEmpty():
		return pushr(l, c)
	// If either branch is a single, glue reduces to pushl/pushr
	case l.isB:
		return pushl(r, c).Pushl(*l.B())
	case r.isB:
		return pushr(l, c).Pushr(*r.B())
	default: // Otherwise, both branches are trees, proceed recursively
		lt := *l.C()
		rt := *r.C()
		ns := nodes(append(append(lt.right, c...), rt.left...))
		nc := glue(lt.child, ns, rt.child)
		return triple(lt.left, rt.right, nc)
	}
}
