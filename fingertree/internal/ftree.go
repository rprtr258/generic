package internal

type ftree struct {
	left, right []any
	child       FingerTree
}

func empty() FingerTree {
	return FingerTree{newUnionTag3A[struct{}, any, *ftree](struct{}{})}
}

func single(a any) FingerTree {
	return FingerTree{newUnionTag3B[struct{}, any, *ftree](a)}
}

func triple(left, right []any, child FingerTree) FingerTree {
	return FingerTree{newUnionTag3C[struct{}, any, *ftree](&ftree{left, right, child})}
}

func (s FingerTree) Pushl(d any) FingerTree {
	return Switch3(
		s.unionTag3,
		func(struct{}) FingerTree { return single(d) },
		func(a any) FingerTree { return triple([]any{d}, []any{a}, empty()) },
		func(t *ftree) FingerTree {
			if len(t.left) < 4 {
				return triple(
					append([]any{d}, t.left...),
					t.right,
					t.child,
				)
			}

			pushdown := node3(
				t.left[1],
				t.left[2],
				t.left[3],
			)

			return triple(
				[]any{d, t.left[0]},
				t.right,
				t.child.Pushl(&pushdown),
			)
		},
	)
}

func (s FingerTree) Popl() (FingerTree, any) {
	var (
		res  FingerTree
		elem any
	)
	s.Switch(
		func(*struct{}) { res, elem = empty(), nil },
		func(a *any) { res, elem = empty(), *a },
		func(f **ftree) {
			t := FingerTree{newUnionTag3C[struct{}, any, *ftree](*f)}
			res, elem = t.Taill(), t.Headl()
		},
	)
	return res, elem
}

func (s FingerTree) Popr() (FingerTree, any) {
	var (
		res  FingerTree
		elem any
	)
	s.Switch(
		func(*struct{}) { res, elem = empty(), nil },
		func(a *any) { res, elem = empty(), *a },
		func(f **ftree) {
			t := FingerTree{newUnionTag3C[struct{}, any, *ftree](*f)}
			res, elem = t.Tailr(), t.Headr()
		},
	)
	return res, elem
}

func (s FingerTree) Pushr(d any) FingerTree {
	return Switch3(
		s.unionTag3,
		func(struct{}) FingerTree { return single(d) },
		func(a any) FingerTree { return triple([]any{a}, []any{d}, empty()) },
		func(t *ftree) FingerTree {
			if len(t.right) < 4 {
				return triple(
					t.left,
					append(t.right, d),
					t.child,
				)
			}

			pushdown := node3(
				t.right[0],
				t.right[1],
				t.right[2],
			)

			return triple(
				t.left,
				[]any{t.right[3], d},
				t.child.Pushr(&pushdown),
			)
		},
	)
}

func (s FingerTree) Iterl(f func(any)) {
	Foldl(s, func(o struct{}, b any) struct{} {
		f(b)
		return o
	}, struct{}{})
}

func (s FingerTree) Iterr(f func(any)) {
	Foldr(s, func(o struct{}, b any) struct{} {
		f(b)
		return o
	}, struct{}{})
}

func (s FingerTree) Headr() any {
	var res any
	s.Switch(
		func(*struct{}) { res = nil },
		func(a *any) { res = *a },
		func(f **ftree) { res = (*f).right[len((*f).right)-1] },
	)
	return res
}

func (s FingerTree) Headl() any {
	var res any
	s.Switch(
		func(*struct{}) { res = nil },
		func(a *any) { res = *a },
		func(f **ftree) { res = (*f).left[0] },
	)
	return res
}

func (s FingerTree) Tailr() FingerTree {
	var res FingerTree
	s.Switch(
		func(*struct{}) { res = empty() },
		func(a *any) { res = empty() },
		func(f **ftree) {
			t := *f
			if right := t.right[:len(t.right)-1]; len(right) > 0 {
				res = triple(t.left, right, t.child)
				return
			}

			if t.child.IsEmpty() {
				tree := FingerTree(empty())
				for _, item := range t.left {
					tree = tree.Pushr(item)
				}
				res = tree
				return
			}

			res = triple(
				t.left,
				t.child.Headr().(*node).ToSlice(),
				t.child.Tailr(),
			)
		},
	)
	return res
}

func (s FingerTree) Taill() FingerTree {
	var res FingerTree
	s.Switch(
		func(*struct{}) { res = empty() },
		func(a *any) { res = empty() },
		func(f **ftree) {
			t := *f
			if left := t.left[1:]; len(left) > 0 {
				res = triple(left, t.right, t.child)
				return
			}

			if t.child.IsEmpty() {
				tree := FingerTree(empty())
				for _, item := range t.right {
					tree = tree.Pushr(item)
				}
				res = tree
				return
			}

			res = triple(
				t.child.Headl().(*node).ToSlice(),
				t.right,
				t.child.Taill(),
			)
		},
	)
	return res
}

func (s FingerTree) IsEmpty() bool {
	var res bool
	s.Switch(
		func(*struct{}) { res = true },
		func(*any) { res = false },
		func(**ftree) { res = false },
	)
	return res
}

func (s FingerTree) Concat(other FingerTree) FingerTree {
	var res FingerTree
	s.Switch(
		func(*struct{}) { res = other },
		func(a *any) { res = other.Pushl(*a) },
		func(f **ftree) {
			res = glue(FingerTree{newUnionTag3C[struct{}, any, *ftree](*f)}, []any{}, other)
		},
	)
	return res
}
