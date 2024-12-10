package cp

import "fmt"

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
