package scripts

import (
	"math/rand"
)

func NewMatrixInt(size int) [][]int {
	randMatrix := make([][]int, size)
	for i := 0; i < size; i++ {
		randMatrix[i] = make([]int, size)
	}
	generateMatrix(randMatrix)

	return randMatrix
}

func generateMatrix(randMatrix [][]int) {
	for i, innerArray := range randMatrix {
		for j := range innerArray {
			item := rand.Intn(100)
			randMatrix[i][j] = item
		}
	}
}
