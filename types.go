package cp

type IntegerNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type DecimalNumber interface {
	~float32 | ~float64
}

type MNode[T any] struct {
	value T
	row   int
	col   int
}

type Matrix[T any] [][]MNode[T]
