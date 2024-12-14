package cp

type Node[T any] struct {
	Value T
	Row   int
	Col   int
}

type Matrix[T any] [][]Node[T]

type MatrixRange struct {
	StartRow, EndRow int
	StartCol, EndCol int
	Rows             int
	Cols             int
}
