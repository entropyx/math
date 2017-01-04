package math

import (
	"math"
	"runtime"
)

// Traspose of a matrix
func Traspose(X [][]float64) (t [][]float64) {
	l1 := len(X[0])
	l2 := len(X)
	for i := 0; i < l1; i++ {
		var row []float64
		for j := 0; j < l2; j++ {
			row = append(row, X[j][i])
		}
		t = append(t, row)
	}
	return
}

// VectorProduct return the scalar product between two vectors.
func VectorProduct(x, y []float64) float64 {
	l := len(x)
	prod := 0.0000
	if l != len(y) {
		panic("The length of vector must be the same!")
	}
	for i := 0; i < l; i++ {
		prod = prod + x[i]*y[i]
	}
	return prod
}

// VectorDiff return the difference between two vectors.
func VectorDiff(x, y []float64) []float64 {
	l1 := len(x)
	l2 := len(y)
	if l1 != l2 {
		panic("The length of vector must be the same!")
	}
	diff := make([]float64, len(x))
	for i := 0; i < l1; i++ {
		diff[i] = x[i] - y[i]
	}
	return diff
}

// VectorAdd return the sum between two vectors.
func VectorAdd(x, y []float64) []float64 {
	l1 := len(x)
	l2 := len(y)
	if l1 != l2 {
		panic("The length of vector must be the same!")
	}
	add := make([]float64, len(x))
	for i := 0; i < l1; i++ {
		add[i] = x[i] + y[i]
	}
	return add
}

// MatrixProduct return the product between two matrix
func MatrixProduct(X, Y [][]float64, c chan [][]float64) {
	var prod [][]float64
	l1 := len(X)
	l2 := len(Y)
	if len(X[0]) != len(Y[0]) {
		panic("Matrix dimension error!")
	}
	for i := 0; i < l1; i++ {
		var row []float64
		for j := 0; j < l2; j++ {
			row = append(row, VectorProduct(X[i], Y[j]))
		}
		prod = append(prod, row)
	}
	c <- prod
}

// ParallelMatrixProd compute distribute the product between two matrix
func ParallelMatrixProd(X, Y [][]float64) (u [][]float64) {
	l1 := len(X)
	if len(X[0]) != len(Y[0]) {
		panic("Matrix dimension error!")
	}
	n := runtime.NumCPU()
	N := int(math.Floor(float64(l1 / n)))
	c := make(chan [][]float64)
	for i := 0; i < n; i++ {
		j1 := i * N
		j2 := (i + 1) * N
		go MatrixProduct(X[j1:j2], Y, c)
	}

	if n*N != l1 {
		go MatrixProduct(X[(n*N):], Y, c)
		n++
	}

	for i := 0; i < n; i++ {
		submatrix := <-c
		u = append(u, submatrix...)
	}
	close(c)
	return u
}
