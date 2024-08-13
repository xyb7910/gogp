package slice

// ToSliceByMapKey is a generic function that converts a map to a slice of its keys.
func ToSliceByMapKey[K comparable, V any](m map[K]V) []K {
	// Create a slice to hold the keys
	keys := make([]K, 0, len(m))

	// Iterate over the map and append each key to the slice
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// ToSliceByMapValue is a generic function that converts a map to a slice of its values.
func ToSliceByMapValue[K comparable, V any](m map[K]V) []V {
	// Create a slice to hold the values
	values := make([]V, 0, len(m))

	// Iterate over the map and append each value to the slice
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

// Map applies a function to each element of a slice and returns a new slice with the results.
func Map[T any, U any](input []T, mapper func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = mapper(v)
	}
	return result
}
