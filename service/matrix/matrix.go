package matrix

import "strconv"

type MatInt [][]int
type MatString [][]string

/*func ConvertItemToFlatFloat(data [][]string) ([]float64, int) {
	var matrix []float64
	var n int
	for _, line := range data {
		n = len(line)
		for _, item := range line {
			floatVar, _ := strconv.ParseFloat(item, 64)
			matrix = append(matrix, floatVar)
		}
	}
	return matrix, n
}

func Multiply(matrixA []float64, matrixB []float64, n int, m int) mat.Dense {
	a := mat.NewDense(n, n, matrixA)
	b := mat.NewDense(m, m, matrixB)
	var mulResult mat.Dense
	mulResult.Mul(a, b)
	return mulResult
}

func ConvertItemToString(result []float64, rows int, cols int) [][]string {
	strResult := make([][]string, cols)
	for i := range strResult {
		strResult[i] = make([]string, rows)
	}
	for i, row := range strResult {
		for j := range row {
			strResult[i][j] = fmt.Sprintf("%.0f", result[i*rows+j])
		}
	}
	return strResult
}*/

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

func MulMatrix(matrix1, matrix2 MatInt) MatInt {
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
