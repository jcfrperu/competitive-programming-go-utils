package cp

type Node[T any] struct {
	Value T
	Row   int
	Col   int
}

type Matrix[T any] [][]Node[T]

func (m Matrix[T]) RowSize() int {
	return len(m)
}

func (m Matrix[T]) ColSize() int {
	return len(m[0])
}

func (m Matrix[T]) MatrixSize() (int, int) {
	return m.RowSize(), m.ColSize()
}

func (m Matrix[T]) IsValidRow(rowIndex int) bool {
	return rowIndex >= 0 && rowIndex < m.RowSize()
}

func (m Matrix[T]) IsValidCol(colIndex int) bool {
	return colIndex >= 0 && colIndex < m.ColSize()
}

func (m Matrix[T]) IsValidRowRange(startRowIndex int, endRowIndex int) bool {
	return m.IsValidRow(startRowIndex) && m.IsValidRow(endRowIndex)
}

func (m Matrix[T]) IsValidColRange(startColIndex int, endColIndex int) bool {
	return m.IsValidCol(startColIndex) && m.IsValidCol(endColIndex)
}

func (m Matrix[T]) RowAt(rowIndex int) []Node[T] {
	return m[rowIndex]
}

func (m Matrix[T]) ColumnAt(colIndex int) []Node[T] {
	rowsNumber := m.ColSize()
	col := make([]Node[T], rowsNumber)
	for i := range col {
		col[i] = m[i][colIndex]
	}
	return col
}

func (m Matrix[T]) Transpose() Matrix[T] {
	rowsNumber := m.ColSize()
	colsNumber := m.RowSize()
	t := make([][]Node[T], rowsNumber)
	for i := range t {
		t[i] = make([]Node[T], colsNumber)
		for j := 0; j < colsNumber; j++ {
			t[i][j] = Node[T]{
				Value: m[j][i].Value,
				Row:   i,
				Col:   j,
			}
		}
	}
	return t
}
