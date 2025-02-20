package slice

func CreateSlice(n int) [][]int{
	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	return matrix
}