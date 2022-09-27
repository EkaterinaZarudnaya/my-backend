package service

import (
	"strconv"
)

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
}*/

func ConvertItemToInt(data [][]string) [][]int {
	matrix := make([][]int, len(data))
	for i, line := range data {
		matrix[i] = make([]int, len(line))
		for j, item := range line {
			intVar, _ := strconv.Atoi(item)
			matrix[i][j] = intVar
		}
	}
	return matrix
}

func MulMatrix(matrix1 [][]int, matrix2 [][]int) [][]int {
	result := make([][]int, len(matrix1))
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

func ConvertItemToString(result [][]int) [][]string {
	strResult := make([][]string, len(result))
	for i, line := range result {
		strResult[i] = make([]string, len(line))
		for j, item := range line {
			strVar := strconv.Itoa(item)
			strResult[i][j] = strVar
		}
	}
	return strResult
}
