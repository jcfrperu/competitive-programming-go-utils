package cp

type Node[T any] struct {
	Value T
	Row   int
	Col   int
}

type Matrix[T any] [][]Node[T]
