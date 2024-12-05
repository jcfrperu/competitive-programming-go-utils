package cp

func BuildMatrix[T int | int64 | float64 | string](input [][]T) Matrix[any] {
	rowSize := len(input)
	colSize := len(input[0])

	matrix := make([][]Node[any], rowSize)
	for i := 0; i < rowSize; i++ {
		matrix[i] = make([]Node[any], colSize)
		for j := 0; j < colSize; j++ {
			matrix[i][j] = Node[any]{
				Value: input[i][j],
				Row:   i,
				Col:   j,
			}
		}
	}
	return matrix
}

func BuildStringMatrix(lines []string) Matrix[string] {
	rowSize := len(lines)
	colSize := len(lines[0])

	matrix := make([][]Node[string], rowSize)
	for i := 0; i < rowSize; i++ {
		matrix[i] = make([]Node[string], colSize)
		for j := 0; j < colSize; j++ {
			matrix[i][j] = Node[string]{
				Value: string(lines[i][j]),
				Row:   i,
				Col:   j,
			}
		}
	}
	return matrix
}
