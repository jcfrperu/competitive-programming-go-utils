package cp

type Node[T any] struct {
	Value T
	Row   int
	Col   int
}

func (n Node[T]) IsValid() bool {
	return n.Value != nil && n.Row >= 0 && n.Col >= 0
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

func (m Matrix[T]) IsValid(rowIndex int, colIndex int) bool {
	return m.IsValidRow(rowIndex) && m.IsValidCol(colIndex)
}

func (m Matrix[T]) IsValidRowRange(startRowIndex int, endRowIndex int) bool {
	return m.IsValidRow(startRowIndex) && m.IsValidRow(endRowIndex) && endRowIndex >= startRowIndex
}

func (m Matrix[T]) IsValidColRange(startColIndex int, endColIndex int) bool {
	return m.IsValidCol(startColIndex) && m.IsValidCol(endColIndex) && endColIndex >= startColIndex
}

func (m Matrix[T]) RowAt(rowIndex int) []Node[T] {
	if !m.IsValidRow(rowIndex) {
		return nil
	}
	return m[rowIndex]
}

func (m Matrix[T]) ColumnAt(colIndex int) []Node[T] {
	if !m.IsValidCol(colIndex) {
		return nil
	}
	rowsNumber := m.ColSize()
	col := make([]Node[T], rowsNumber)
	for i := range col {
		col[i] = m[i][colIndex]
	}
	return col
}

func (m Matrix[T]) HasValidNeighbors(rowIndex int, colIndex int) bool {
	return m.GetUp(rowIndex, colIndex).IsValid() &&
		m.GetRight(rowIndex, colIndex).IsValid() &&
		m.GetDown(rowIndex, colIndex).IsValid() &&
		m.GetLeft(rowIndex, colIndex).IsValid()
}

func (m Matrix[T]) HasValidCrossNeighbors(rowIndex int, colIndex int) bool {
	return m.GetUpLeft(rowIndex, colIndex).IsValid() &&
		m.GetUpRight(rowIndex, colIndex).IsValid() &&
		m.GetDownRight(rowIndex, colIndex).IsValid() &&
		m.GetDownLeft(rowIndex, colIndex).IsValid()
}

func (m Matrix[T]) HasValidAllNeighbors(rowIndex int, colIndex int) bool {
	return m.HasValidNeighbors(rowIndex, colIndex) && m.HasValidCrossNeighbors(rowIndex, colIndex)
}

func (m Matrix[T]) GetNeighbors(rowIndex int, colIndex int) []Node[T] {
	neighbors := make([]Node[T], 0)
	neighbors = append(neighbors, m.GetUp(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetRight(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDown(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetLeft(rowIndex, colIndex))
	return neighbors
}

func (m Matrix[T]) GetAllCrossNeighbors(rowIndex int, colIndex int) []Node[T] {
	neighbors := make([]Node[T], 0)
	neighbors = append(neighbors, m.GetUpLeft(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetUpRight(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownRight(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownLeft(rowIndex, colIndex))
	return neighbors
}

func (m Matrix[T]) GetAllNeighbors(rowIndex int, colIndex int) []Node[T] {
	neighbors := make([]Node[T], 0)
	neighbors = append(neighbors, m.GetUpLeft(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetUp(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetUpRight(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetRight(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownRight(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDown(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownLeft(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetLeft(rowIndex, colIndex))
	return neighbors
}

func (m Matrix[T]) GetAt(rowIndex int, colIndex int) Node[T] {
	if !m.IsValid(rowIndex, colIndex) {
		return Node[T]{}
	}
	return m[rowIndex][colIndex]
}

func (m Matrix[T]) GetLeft(rowIndex int, colIndex int) Node[T] {
	if !m.IsValid(rowIndex, colIndex) || !m.IsValid(rowIndex, colIndex-1) {
		return Node[T]{}
	}
	return m[rowIndex][colIndex-1]
}

func (m Matrix[T]) GetRight(rowIndex int, colIndex int) Node[T] {
	if !m.IsValid(rowIndex, colIndex) || !m.IsValid(rowIndex, colIndex+1) {
		return Node[T]{}
	}
	return m[rowIndex][colIndex+1]
}

func (m Matrix[T]) GetUp(rowIndex int, colIndex int) Node[T] {
	if !m.IsValid(rowIndex, colIndex) || !m.IsValid(rowIndex-1, colIndex) {
		return Node[T]{}
	}
	return m[rowIndex-1][colIndex]
}

func (m Matrix[T]) GetDown(rowIndex int, colIndex int) Node[T] {
	if !m.IsValid(rowIndex, colIndex) || !m.IsValid(rowIndex+1, colIndex) {
		return Node[T]{}
	}
	return m[rowIndex+1][colIndex]
}

func (m Matrix[T]) GetUpLeft(rowIndex int, colIndex int) Node[T] {
	if !m.IsValid(rowIndex, colIndex) || !m.IsValid(rowIndex-1, colIndex-1) {
		return Node[T]{}
	}
	return m[rowIndex-1][colIndex-1]
}

func (m Matrix[T]) GetUpRight(rowIndex int, colIndex int) Node[T] {
	if !m.IsValid(rowIndex, colIndex) || !m.IsValid(rowIndex-1, colIndex+1) {
		return Node[T]{}
	}
	return m[rowIndex-1][colIndex+1]
}

func (m Matrix[T]) GetDownLeft(rowIndex int, colIndex int) Node[T] {
	if !m.IsValid(rowIndex, colIndex) || !m.IsValid(rowIndex+1, colIndex-1) {
		return Node[T]{}
	}
	return m[rowIndex+1][colIndex-1]
}

func (m Matrix[T]) GetDownRight(rowIndex int, colIndex int) Node[T] {
	if !m.IsValid(rowIndex, colIndex) || !m.IsValid(rowIndex+1, colIndex+1) {
		return Node[T]{}
	}
	return m[rowIndex+1][colIndex+1]
}

func (m Matrix[T]) Transpose() Matrix[T] {
	rowSize := m.ColSize()
	colSize := m.RowSize()

	t := make([][]Node[T], rowSize)
	for i := range t {
		t[i] = make([]Node[T], colSize)
		for j := 0; j < colSize; j++ {
			t[i][j] = Node[T]{
				Value: m[j][i].Value,
				Row:   i,
				Col:   j,
			}
		}
	}
	return t
}
