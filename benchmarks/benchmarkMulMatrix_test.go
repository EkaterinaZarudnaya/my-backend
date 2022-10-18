package benchmarks

import (
	"my-backend/service/matrix"
	"testing"
)

func BenchmarkMulMatrix(b *testing.B) {
	//
	for n := 0; n < b.N; n++ {
		matrix.MulMatrix(matrixA, matrixB)
	}
}
