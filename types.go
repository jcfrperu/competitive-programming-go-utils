package cp

type Cell[T any] struct {
	Value T
	Row   int
	Col   int
}

type Matrix[T any] [][]Cell[T]

type MatrixRange struct {
	StartRow, EndRow int
	StartCol, EndCol int
	Rows             int
	Cols             int
}

type ListRange struct {
	Start int
	End   int
	Size  int
}
