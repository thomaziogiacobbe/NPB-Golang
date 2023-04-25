package commons

func CreateMatrix(rows int, columns int) [][]int64 {
	a := make([]int64, columns*rows)
	m := make([][]int64, rows)
	lo, hi := 0, columns
	for i := range m {
		m[i] = a[lo:hi:hi]
		lo, hi = hi, hi+columns
	}
	return m
}

func ResetMatrix(matrix [][]int64) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = 0
		}
	}
}
