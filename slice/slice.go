package slice

// Reduce reduces a slice to a single value using a specified reduction function.
func Reduce[T any, U any](input []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, v := range input {
		result = reducer(result, v)
	}
	return result
}
