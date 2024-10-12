package internal

type node struct {
	unionTag[[2]any, [3]any]
}

func node2(a, b any) node {
	return node{NewUnionTagA[[2]any, [3]any]([2]any{a, b})}
}

func node3(a, b, c any) node {
	return node{NewUnionTagB[[2]any, [3]any]([3]any{a, b, c})}
}

func (n node) ToSlice() []any {
	var res []any
	n.Switch(
		func(a *[2]any) { res = (*a)[:] },
		func(a *[3]any) { res = (*a)[:] },
	)
	return res
}

func nodeFoldl[R any](n node, f func(R, any) R, initial R) R {
	var res R
	n.Switch(
		func(a *[2]any) {
			res = f(f(initial, a[0]), a[1])
		},
		func(a *[3]any) {
			res = f(f(f(initial, a[0]), a[1]), a[2])
		},
	)
	return res
}

func nodeFoldr[R any](n node, f func(R, any) R, initial R) R {
	var res R
	n.Switch(
		func(a *[2]any) {
			res = f(f(initial, a[1]), a[0])
		},
		func(a *[3]any) {
			res = f(f(f(initial, a[2]), a[1]), a[0])
		},
	)
	return res
}

func (n node) Iterl(f func(any)) {
	n.Switch(
		func(a *[2]any) {
			f(a[0])
			f(a[1])
		},
		func(a *[3]any) {
			f(a[0])
			f(a[1])
			f(a[2])
		},
	)
}

func (n node) Iterr(f func(any)) {
	n.Switch(
		func(a *[2]any) {
			f(a[1])
			f(a[0])
		},
		func(a *[3]any) {
			f(a[2])
			f(a[1])
			f(a[0])
		},
	)
}
