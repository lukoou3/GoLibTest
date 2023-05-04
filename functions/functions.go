package functions

func filterImpl[A any, T ~[]A](slice T, p func(A) bool, isFlipped bool) T {
	rst := make(T, 0)
	for _, e := range slice {
		if p(e) != isFlipped {
			rst = append(rst, e)
		}
	}
	return rst
}

func Filter[A any, T ~[]A](slice T, p func(A) bool) T {
	return filterImpl(slice, p, false)
}

func FilterNot[A any, T ~[]A](slice T, p func(A) bool) T {
	return filterImpl(slice, p, true)
}

func Map[A any, T ~[]A, B any](slice T, f func(A) B) []B {
	rst := make([]B, 0, len(slice))
	for _, x := range slice {
		rst = append(rst, f(x))
	}
	return rst
}

func ForEach[A any](slice []A, f func(A)) {
	for _, x := range slice {
		f(x)
	}
}
