package mapx

// ToMap converts a slice to a map.
// convert []comparable to map[comparable]struct{}, key is comparable, value is struct{}{}
func ToMap[T comparable](src []T) map[T]struct{} {
	res := make(map[T]struct{}, len(src))
	for _, v := range src {
		res[v] = struct{}{}
	}
	return res
}
