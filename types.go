package gogp

// RealNumber 实数类型
type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

// Number 数字类型
type Number interface {
	RealNumber | ~complex64 | ~complex128
}
