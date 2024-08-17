package slicex

// Shrink shrinks the cap of the slicex.
func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	n, changed := calCap(c, l)
	if !changed {
		return src
	}
	s := make([]T, 0, n)
	s = append(s, src...)
	return s
}

// calCap returns the new cap and whether it is changed.
func calCap(c, l int) (int, bool) {
	if c < 64 {
		return c, false
	}
	if c > 2048 && (c/l >= 2) {
		factor := 0.625
		return int(float32(c) * float32(factor)), true
	}
	if c > 2048 && (c/l >= 4) {
		return c / 2, true
	}
	return c, false
}
