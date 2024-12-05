package cp

type MNode[T any] struct {
	Value T
	Row   int
	Col   int
}

type Matrix[T any] [][]MNode[T]
