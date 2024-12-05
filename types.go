package cp

type Node[T any] struct {
	Value T
	Row   int
	Col   int
}

type Matrix[T any] struct {
	Matrix  [][]Node[T]
	RowSize int
	ColSize int
}

type DataMatrix[T any] struct {
	Matrix           Matrix[T]
	TransposedMatrix Matrix[T]
}
