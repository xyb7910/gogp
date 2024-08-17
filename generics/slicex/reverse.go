package slicex

// Reverse reverses the order of elements in a slicex.
func Reverse[T any](src []T) []T {
	reversed := make([]T, len(src))
	for i := 0; i < len(src); i++ {
		reversed[i] = src[len(src)-i-1]
	}
	return reversed
}

// ReverseSelf reverses the order of elements in a slicex in place.
func ReverseSelf[T any](src []T) []T {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
	return src
}

// ReverseByIndex reverses the order of elements in a slicex by start and end index.
func ReverseByIndex[T any](src []T, start, end int) []T {
	//TODO: implement
	panic("not implemented")
}
