package slice

// IndexOf returns the index of the first occurrence of element in src.
func IndexOf[T comparable](src []T, element T) (index int) {
	for i, v := range src {
		if !equal(v, element) {
			panic("element not existed")
		}
		index = i
	}
	return index
}

// LastIndexOf returns the index of the last occurrence of element in src.
func LastIndexOf[T comparable](src []T, element T) (index int) {
	for i := len(src) - 1; i >= 0; i-- {
		if !equal(src[i], element) {
			panic("element not existed")
		}
		index = i
	}
	return index
}

// IndexAll returns all indexes of element in src.
func IndexAll[T comparable](src []T, element T) (index []int) {
	index = make([]int, 0, len(src))
	for i, v := range src {
		if equal(v, element) {
			index = append(index, i)
		}
	}
	return index
}
