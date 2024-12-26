package cp

import "fmt"

// methods for Cell type

func (n Cell[T]) IsValid() bool {
	return n.Row >= 0 && n.Col >= 0
}

func (n Cell[T]) Update(value T) Cell[T] {
	n.Value = value
	return n
}

// methods for Matrix range

func (r MatrixRange) NewMatrixRangeFromSizes(rows int, cols int) MatrixRange {
	return r.NewMatrixRange(0, rows-1, 0, cols-1)
}

func (r MatrixRange) NewMatrixRange(startRow int, endRow int, startCol int, endCol int) MatrixRange {
	return MatrixRange{
		StartRow: startRow,
		EndRow:   endRow,
		StartCol: startCol,
		EndCol:   endCol,
		Rows:     endRow - startRow + 1,
		Cols:     endCol - startCol + 1,
	}
}

func (r MatrixRange) Inside(rowIndex int, colIndex int) bool {
	return rowIndex >= r.StartRow && rowIndex <= r.EndRow && colIndex >= r.StartCol && colIndex <= r.EndCol
}

func (r MatrixRange) InsideRow(rowIndex int) bool {
	return rowIndex >= r.StartRow && rowIndex <= r.EndRow
}

func (r MatrixRange) InsideStartRow(rowIndex int) bool {
	return rowIndex >= r.StartRow
}

func (r MatrixRange) InsideEndRow(rowIndex int) bool {
	return rowIndex <= r.EndRow
}

func (r MatrixRange) InsideCol(colIndex int) bool {
	return colIndex >= r.StartCol && colIndex <= r.EndCol
}

func (r MatrixRange) InsideStartCol(colIndex int) bool {
	return colIndex >= r.StartCol
}

func (r MatrixRange) InsideEndCol(colIndex int) bool {
	return colIndex <= r.EndCol
}

// methods for Matrix type

func (m Matrix[T]) Rows() int {
	return len(m)
}

func (m Matrix[T]) Cols() int {
	return len(m[0])
}

func (m Matrix[T]) HasRowAt(rowIndex int) bool {
	return rowIndex >= 0 && rowIndex < m.Rows()
}

func (m Matrix[T]) HasColAt(colIndex int) bool {
	return colIndex >= 0 && colIndex < m.Cols()
}

func (m Matrix[T]) IsValidAt(rowIndex int, colIndex int) bool {
	return m.HasRowAt(rowIndex) && m.HasColAt(colIndex)
}

func (m Matrix[T]) SetAt(rowIndex int, colIndex int, newValue T) {
	m[rowIndex][colIndex].Value = newValue
}

func (m Matrix[T]) GetAt(rowIndex int, colIndex int) T {
	return m[rowIndex][colIndex].Value
}

func (m Matrix[T]) SetNodeAt(node Cell[T]) {
	m[node.Row][node.Col] = node
}

func (m Matrix[T]) GetNodeAt(node Cell[T]) Cell[T] {
	return m[node.Row][node.Col]
}

func (m Matrix[T]) GetRowAt(rowIndex int) []Cell[T] {
	return m[rowIndex]
}

func (m Matrix[T]) GetColAt(colIndex int) []Cell[T] {
	rowsNumber := m.Cols()
	col := make([]Cell[T], rowsNumber)
	for i := range col {
		col[i] = m[i][colIndex]
	}
	return col
}

func (m Matrix[T]) GetCrossNeighborsAt(rowIndex int, colIndex int) []Cell[T] {
	neighbors := make([]Cell[T], 0)
	neighbors = append(neighbors, m.GetLeftAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetUpAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownAt(rowIndex, colIndex))
	return neighbors
}

func (m Matrix[T]) GetDiagonalNeighborsAt(rowIndex int, colIndex int) []Cell[T] {
	neighbors := make([]Cell[T], 0)
	neighbors = append(neighbors, m.GetUpLeftAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetUpRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownLeftAt(rowIndex, colIndex))
	return neighbors
}

func (m Matrix[T]) GetNeighborsAt(rowIndex int, colIndex int) []Cell[T] {
	neighbors := make([]Cell[T], 0)
	neighbors = append(neighbors, m.GetLeftAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetUpLeftAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetUpAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetUpRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownLeftAt(rowIndex, colIndex))
	return neighbors
}

func (m Matrix[T]) GetValidCrossNeighborsAt(rowIndex int, colIndex int) []Cell[T] {
	neighbors := m.GetCrossNeighborsAt(rowIndex, colIndex)
	validNodes := make([]Cell[T], 0)
	for _, node := range neighbors {
		if node.IsValid() {
			validNodes = append(validNodes, node)
		}
	}
	return validNodes
}

func (m Matrix[T]) GetValidDiagonalNeighborsAt(rowIndex int, colIndex int) []Cell[T] {
	neighbors := m.GetDiagonalNeighborsAt(rowIndex, colIndex)
	validNodes := make([]Cell[T], 0)
	for _, node := range neighbors {
		if node.IsValid() {
			validNodes = append(validNodes, node)
		}
	}
	return validNodes
}

func (m Matrix[T]) GetValidNeighborsAt(rowIndex int, colIndex int) []Cell[T] {
	neighbors := m.GetNeighborsAt(rowIndex, colIndex)
	validNodes := make([]Cell[T], 0)
	for _, node := range neighbors {
		if node.IsValid() {
			validNodes = append(validNodes, node)
		}
	}
	return validNodes
}

func (m Matrix[T]) GetLeftAt(rowIndex int, colIndex int) Cell[T] {
	if !m.IsValidAt(rowIndex, colIndex) || !m.IsValidAt(rowIndex, colIndex-1) {
		return Cell[T]{Row: -1, Col: -1}
	}
	return m[rowIndex][colIndex-1]
}

func (m Matrix[T]) GetRightAt(rowIndex int, colIndex int) Cell[T] {
	if !m.IsValidAt(rowIndex, colIndex) || !m.IsValidAt(rowIndex, colIndex+1) {
		return Cell[T]{Row: -1, Col: -1}
	}
	return m[rowIndex][colIndex+1]
}

func (m Matrix[T]) GetUpAt(rowIndex int, colIndex int) Cell[T] {
	if !m.IsValidAt(rowIndex, colIndex) || !m.IsValidAt(rowIndex-1, colIndex) {
		return Cell[T]{Row: -1, Col: -1}
	}
	return m[rowIndex-1][colIndex]
}

func (m Matrix[T]) GetDownAt(rowIndex int, colIndex int) Cell[T] {
	if !m.IsValidAt(rowIndex, colIndex) || !m.IsValidAt(rowIndex+1, colIndex) {
		return Cell[T]{Row: -1, Col: -1}
	}
	return m[rowIndex+1][colIndex]
}

func (m Matrix[T]) GetUpLeftAt(rowIndex int, colIndex int) Cell[T] {
	if !m.IsValidAt(rowIndex, colIndex) || !m.IsValidAt(rowIndex-1, colIndex-1) {
		return Cell[T]{Row: -1, Col: -1}
	}
	return m[rowIndex-1][colIndex-1]
}

func (m Matrix[T]) GetUpRightAt(rowIndex int, colIndex int) Cell[T] {
	if !m.IsValidAt(rowIndex, colIndex) || !m.IsValidAt(rowIndex-1, colIndex+1) {
		return Cell[T]{Row: -1, Col: -1}
	}
	return m[rowIndex-1][colIndex+1]
}

func (m Matrix[T]) GetDownLeftAt(rowIndex int, colIndex int) Cell[T] {
	if !m.IsValidAt(rowIndex, colIndex) || !m.IsValidAt(rowIndex+1, colIndex-1) {
		return Cell[T]{Row: -1, Col: -1}
	}
	return m[rowIndex+1][colIndex-1]
}

func (m Matrix[T]) GetDownRightAt(rowIndex int, colIndex int) Cell[T] {
	if !m.IsValidAt(rowIndex, colIndex) || !m.IsValidAt(rowIndex+1, colIndex+1) {
		return Cell[T]{Row: -1, Col: -1}
	}
	return m[rowIndex+1][colIndex+1]
}

func (m Matrix[T]) GetCrossNeighbors(node Cell[T]) []Cell[T] {
	return m.GetCrossNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetDiagonalNeighbors(node Cell[T]) []Cell[T] {
	return m.GetDiagonalNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetNeighbors(node Cell[T]) []Cell[T] {
	return m.GetNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetNeighborsWith(pattern string, node Cell[T]) []Cell[T] {
	rowIndex := node.Row
	colIndex := node.Col

	neighbors := make([]Cell[T], 0)

	// cross and diagonal
	if pattern == "+" {
		neighbors = append(neighbors, m.GetLeftAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetUpAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetRightAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetDownAt(rowIndex, colIndex))
		return neighbors

	}
	if pattern == "x" {
		neighbors = append(neighbors, m.GetUpLeftAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetUpRightAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetDownRightAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetDownLeftAt(rowIndex, colIndex))
		return neighbors
	}
	// lateral al vertical
	if pattern == "-" {
		neighbors = append(neighbors, m.GetLeftAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetRightAt(rowIndex, colIndex))
		return neighbors
	}
	if pattern == "|" {
		neighbors = append(neighbors, m.GetUpAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetDownAt(rowIndex, colIndex))
		return neighbors
	}
	// single diagonals
	if pattern == "/" {
		neighbors = append(neighbors, m.GetDownLeftAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetUpRightAt(rowIndex, colIndex))
		return neighbors
	}
	if pattern == "\\" {
		neighbors = append(neighbors, m.GetUpLeftAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetDownRightAt(rowIndex, colIndex))
		return neighbors
	}
	// others
	if pattern == "->" {
		neighbors = append(neighbors, m.GetLeftAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetRightAt(rowIndex, colIndex))
		return neighbors
	}
	if pattern == "<-" {
		neighbors = append(neighbors, m.GetRightAt(rowIndex, colIndex))
		neighbors = append(neighbors, m.GetLeftAt(rowIndex, colIndex))
		return neighbors
	}

	// by default returns all neighbors
	return m.GetNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetValidCrossNeighbors(node Cell[T]) []Cell[T] {
	return m.GetValidCrossNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetValidDiagonalNeighbors(node Cell[T]) []Cell[T] {
	return m.GetValidDiagonalNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetValidNeighbors(node Cell[T]) []Cell[T] {
	return m.GetValidNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) IsValid(node Cell[T]) bool {
	return m.IsValidAt(node.Row, node.Col)
}

func (m Matrix[T]) GetLeft(node Cell[T]) Cell[T] {
	return m.GetLeftAt(node.Row, node.Col)
}

func (m Matrix[T]) GetRight(node Cell[T]) Cell[T] {
	return m.GetRightAt(node.Row, node.Col)
}

func (m Matrix[T]) GetUp(node Cell[T]) Cell[T] {
	return m.GetUpAt(node.Row, node.Col)
}

func (m Matrix[T]) GetDown(node Cell[T]) Cell[T] {
	return m.GetDownAt(node.Row, node.Col)
}

func (m Matrix[T]) GetUpLeft(node Cell[T]) Cell[T] {
	return m.GetUpLeftAt(node.Row, node.Col)
}

func (m Matrix[T]) GetUpRight(node Cell[T]) Cell[T] {
	return m.GetUpRightAt(node.Row, node.Col)
}

func (m Matrix[T]) GetDownLeft(node Cell[T]) Cell[T] {
	return m.GetDownLeftAt(node.Row, node.Col)
}

func (m Matrix[T]) GetDownRight(node Cell[T]) Cell[T] {
	return m.GetDownRightAt(node.Row, node.Col)
}

func (m Matrix[T]) Transpose() Matrix[T] {
	rows := m.Cols()
	cols := m.Rows()
	t := make([][]Cell[T], rows)
	for i := range t {
		t[i] = make([]Cell[T], cols)
		for j := 0; j < cols; j++ {
			t[i][j] = Cell[T]{
				Value: m[j][i].Value,
				Row:   i,
				Col:   j,
			}
		}
	}
	return t
}

func (m Matrix[T]) GetRange() MatrixRange {
	return MatrixRange{
		StartRow: 0,
		EndRow:   m.Rows() - 1,
		StartCol: 0,
		EndCol:   m.Cols() - 1,
		Rows:     m.Rows(),
		Cols:     m.Cols(),
	}
}

// other functions

func BuildMatrix[T any](input [][]T) Matrix[T] {
	rows := len(input)
	cols := len(input[0])
	matrix := make([][]Cell[T], rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]Cell[T], cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = Cell[T]{
				Value: input[i][j],
				Row:   i,
				Col:   j,
			}
		}
	}
	return matrix
}

func PrintMatrix[T any](m Matrix[T]) {
	rows := m.Rows()
	cols := m.Cols()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%v", m[i][j].Value)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func BuildStrMatrix(lines []string) Matrix[string] {
	rows := len(lines)
	cols := len(lines[0])
	matrix := make([][]Cell[string], rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]Cell[string], cols)
		for j := 0; j < cols; j++ {
			value := string(lines[i][j])
			matrix[i][j] = Cell[string]{
				Value: value,
				Row:   i,
				Col:   j,
			}
		}
	}
	return matrix
}

func PrintStrMatrix(m Matrix[string]) {
	rows := m.Rows()
	cols := m.Cols()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%s", m[i][j].Value)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
