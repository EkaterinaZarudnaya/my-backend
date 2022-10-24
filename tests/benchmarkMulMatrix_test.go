package tests

import (
	"fmt"
	"my-backend/scripts"
	"my-backend/service/matrix"
	"testing"
)

func BenchmarkMulMatrix(b *testing.B) {
	matrixAB := scripts.NewMatrixInt(1000)
	b.Run(fmt.Sprintf("Matrix multiplication."), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			matrix.MulMatrix(matrixAB, matrixAB)
		}
	})
}
