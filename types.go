package cp

type MNode[T any] struct {
	value T
	row   int
	col   int
}

type Matrix[T any] [][]MNode[T]
