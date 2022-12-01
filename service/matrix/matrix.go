package matrix

import (
	"strconv"
	"sync"
)

type MatInt [][]int
type MatString [][]string

func (ms MatString) ToInt() MatInt {
	matrix := make(MatInt, len(ms))
	for i, line := range ms {
		matrix[i] = make([]int, len(line))
		for j, item := range line {
			intVar, _ := strconv.Atoi(item)
			matrix[i][j] = intVar
		}
	}
	return matrix
}

/*func MulMatrix(matrix1, matrix2 MatInt) MatInt {
	result := make(MatInt, len(matrix1))
	for i := 0; i < len(matrix1); i++ {
		result[i] = make([]int, len(matrix1))
		for j := 0; j < len(matrix2); j++ {
			for k := 0; k < len(matrix2); k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return result
}*/

func MulMatrix(matrix1, matrix2 MatInt) MatInt {
	var wg sync.WaitGroup
	result := make(MatInt, len(matrix1))

	for i := 0; i < len(matrix1); i++ {
		result[i] = make([]int, len(matrix1))
		for j := 0; j < len(matrix2); j++ {
			wg.Add(1)
			go func(i, j int, matrix1, matrix2, result MatInt) {
				wg.Done()
				for k := 0; k < len(matrix2); k++ {
					result[i][j] += matrix1[i][k] * matrix2[k][j]
				}
			}(i, j, matrix1, matrix2, result)
		}
	}
	return result
}

func (m MatInt) ToString() [][]string {
	strResult := make([][]string, len(m))
	for i, line := range m {
		strResult[i] = make([]string, len(line))
		for j, item := range line {
			strVar := strconv.Itoa(item)
			strResult[i][j] = strVar
		}
	}
	return strResult
}
