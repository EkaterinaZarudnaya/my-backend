package matrix

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"strconv"
)

func ConvertItemToFlatFloat(data [][]string) ([]float64, int) {
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
}
