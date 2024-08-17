package slicex

// Find returns the first element in the slicex that matches the match function.
// If no element matches, it returns the zero value of the slicex element type and false.
func Find[T any](src []T, match matchFunc[T]) (T, bool) {
	for _, val := range src {
		if match(val) {
			return val, true
		}
	}
	var t T
	return t, false
}

// FindAll returns all elements in the slicex that matches the match function.
// If no element matches, it returns an empty slicex.
func FindAll[T any](src []T, match matchFunc[T]) []T {
	res := make([]T, 0, len(src)>>3+1)
	for _, val := range src {
		if match(val) {
			res = append(res, val)
		}
	}
	return res
}
