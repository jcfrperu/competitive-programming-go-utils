package cp

import "fmt"

// methods for Node type

func (n Node[T]) IsValid() bool {
	return n.Row >= 0 && n.Col >= 0
}

func (n Node[T]) Update(value T) Node[T] {
	n.Value = value
	return n
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

func (m Matrix[T]) HasNodeAt(rowIndex int, colIndex int) bool {
	return m.HasRowAt(rowIndex) && m.HasColAt(colIndex)
}

func (m Matrix[T]) HasRowRangeAt(startRowIndex int, endRowIndex int) bool {
	return m.HasRowAt(startRowIndex) && m.HasRowAt(endRowIndex) && endRowIndex >= startRowIndex
}

func (m Matrix[T]) HasColRangeAt(startColIndex int, endColIndex int) bool {
	return m.HasColAt(startColIndex) && m.HasColAt(endColIndex) && endColIndex >= startColIndex
}

func (m Matrix[T]) GetRowAt(rowIndex int) []Node[T] {
	return m[rowIndex]
}

func (m Matrix[T]) GetColAt(colIndex int) []Node[T] {
	rowsNumber := m.Cols()
	col := make([]Node[T], rowsNumber)
	for i := range col {
		col[i] = m[i][colIndex]
	}
	return col
}

func (m Matrix[T]) GetCrossNeighborsAt(rowIndex int, colIndex int) []Node[T] {
	neighbors := make([]Node[T], 0)
	neighbors = append(neighbors, m.GetUpAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetLeftAt(rowIndex, colIndex))
	return neighbors
}

func (m Matrix[T]) GetDiagonalNeighborsAt(rowIndex int, colIndex int) []Node[T] {
	neighbors := make([]Node[T], 0)
	neighbors = append(neighbors, m.GetUpLeftAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetUpRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownLeftAt(rowIndex, colIndex))
	return neighbors
}

func (m Matrix[T]) GetNeighborsAt(rowIndex int, colIndex int) []Node[T] {
	neighbors := make([]Node[T], 0)
	neighbors = append(neighbors, m.GetUpLeftAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetUpAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetUpRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownRightAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetDownLeftAt(rowIndex, colIndex))
	neighbors = append(neighbors, m.GetLeftAt(rowIndex, colIndex))
	return neighbors
}

func (m Matrix[T]) GetValidCrossNeighborsAt(rowIndex int, colIndex int) []Node[T] {
	neighbors := m.GetCrossNeighborsAt(rowIndex, colIndex)
	validNodes := make([]Node[T], 0)
	for _, node := range neighbors {
		if node.IsValid() {
			validNodes = append(validNodes, node)
		}
	}
	return validNodes
}

func (m Matrix[T]) GetValidDiagonalNeighborsAt(rowIndex int, colIndex int) []Node[T] {
	neighbors := m.GetDiagonalNeighborsAt(rowIndex, colIndex)
	validNodes := make([]Node[T], 0)
	for _, node := range neighbors {
		if node.IsValid() {
			validNodes = append(validNodes, node)
		}
	}
	return validNodes
}

func (m Matrix[T]) GetValidNeighborsAt(rowIndex int, colIndex int) []Node[T] {
	neighbors := m.GetNeighborsAt(rowIndex, colIndex)
	validNodes := make([]Node[T], 0)
	for _, node := range neighbors {
		if node.IsValid() {
			validNodes = append(validNodes, node)
		}
	}
	return validNodes
}

func (m Matrix[T]) GetLeftAt(rowIndex int, colIndex int) Node[T] {
	if !m.HasNodeAt(rowIndex, colIndex) || !m.HasNodeAt(rowIndex, colIndex-1) {
		return Node[T]{Row: -1, Col: -1}
	}
	return m[rowIndex][colIndex-1]
}

func (m Matrix[T]) GetRightAt(rowIndex int, colIndex int) Node[T] {
	if !m.HasNodeAt(rowIndex, colIndex) || !m.HasNodeAt(rowIndex, colIndex+1) {
		return Node[T]{Row: -1, Col: -1}
	}
	return m[rowIndex][colIndex+1]
}

func (m Matrix[T]) GetUpAt(rowIndex int, colIndex int) Node[T] {
	if !m.HasNodeAt(rowIndex, colIndex) || !m.HasNodeAt(rowIndex-1, colIndex) {
		return Node[T]{Row: -1, Col: -1}
	}
	return m[rowIndex-1][colIndex]
}

func (m Matrix[T]) GetDownAt(rowIndex int, colIndex int) Node[T] {
	if !m.HasNodeAt(rowIndex, colIndex) || !m.HasNodeAt(rowIndex+1, colIndex) {
		return Node[T]{Row: -1, Col: -1}
	}
	return m[rowIndex+1][colIndex]
}

func (m Matrix[T]) GetUpLeftAt(rowIndex int, colIndex int) Node[T] {
	if !m.HasNodeAt(rowIndex, colIndex) || !m.HasNodeAt(rowIndex-1, colIndex-1) {
		return Node[T]{Row: -1, Col: -1}
	}
	return m[rowIndex-1][colIndex-1]
}

func (m Matrix[T]) GetUpRightAt(rowIndex int, colIndex int) Node[T] {
	if !m.HasNodeAt(rowIndex, colIndex) || !m.HasNodeAt(rowIndex-1, colIndex+1) {
		return Node[T]{Row: -1, Col: -1}
	}
	return m[rowIndex-1][colIndex+1]
}

func (m Matrix[T]) GetDownLeftAt(rowIndex int, colIndex int) Node[T] {
	if !m.HasNodeAt(rowIndex, colIndex) || !m.HasNodeAt(rowIndex+1, colIndex-1) {
		return Node[T]{Row: -1, Col: -1}
	}
	return m[rowIndex+1][colIndex-1]
}

func (m Matrix[T]) GetDownRightAt(rowIndex int, colIndex int) Node[T] {
	if !m.HasNodeAt(rowIndex, colIndex) || !m.HasNodeAt(rowIndex+1, colIndex+1) {
		return Node[T]{Row: -1, Col: -1}
	}
	return m[rowIndex+1][colIndex+1]
}

func (m Matrix[T]) GetCrossNeighbors(node Node[T]) []Node[T] {
	return m.GetCrossNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetDiagonalNeighbors(node Node[T]) []Node[T] {
	return m.GetDiagonalNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetNeighbors(node Node[T]) []Node[T] {
	return m.GetNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetValidCrossNeighbors(node Node[T]) []Node[T] {
	return m.GetValidCrossNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetValidDiagonalNeighbors(node Node[T]) []Node[T] {
	return m.GetValidDiagonalNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetValidNeighbors(node Node[T]) []Node[T] {
	return m.GetValidNeighborsAt(node.Row, node.Col)
}

func (m Matrix[T]) GetLeft(node Node[T]) Node[T] {
	return m.GetLeftAt(node.Row, node.Col)
}

func (m Matrix[T]) GetRight(node Node[T]) Node[T] {
	return m.GetRightAt(node.Row, node.Col)
}

func (m Matrix[T]) GetUp(node Node[T]) Node[T] {
	return m.GetUpAt(node.Row, node.Col)
}

func (m Matrix[T]) GetDown(node Node[T]) Node[T] {
	return m.GetDownAt(node.Row, node.Col)
}

func (m Matrix[T]) GetUpLeft(node Node[T]) Node[T] {
	return m.GetUpLeftAt(node.Row, node.Col)
}

func (m Matrix[T]) GetUpRight(node Node[T]) Node[T] {
	return m.GetUpRightAt(node.Row, node.Col)
}

func (m Matrix[T]) GetDownLeft(node Node[T]) Node[T] {
	return m.GetDownLeftAt(node.Row, node.Col)
}

func (m Matrix[T]) GetDownRight(node Node[T]) Node[T] {
	return m.GetDownRightAt(node.Row, node.Col)
}

func (m Matrix[T]) Transpose() Matrix[T] {
	rowSize := m.Cols()
	colSize := m.Rows()

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

// other functions

func BuildMatrix[T any](input [][]T) Matrix[T] {
	rows := len(input)
	cols := len(input[0])
	matrix := make([][]Node[T], rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]Node[T], cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = Node[T]{
				Value: input[i][j],
				Row:   i,
				Col:   j,
			}
		}
	}
	return matrix
}

func BuildStringMatrix(lines []string) Matrix[string] {
	rows := len(lines)
	cols := len(lines[0])

	matrix := make([][]Node[string], rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]Node[string], cols)
		for j := 0; j < cols; j++ {
			value := string(lines[i][j])
			matrix[i][j] = Node[string]{
				Value: value,
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
}

func PrintStringMatrix(m Matrix[string]) {
	rows := m.Rows()
	cols := m.Cols()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%s", m[i][j].Value)
		}
		fmt.Printf("\n")
	}
}
