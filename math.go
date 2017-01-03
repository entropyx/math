package math

import (
	"math"
	"runtime"
)

// Vectorprod return the scalar product between two vectors.
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

func MatrixProduct(X, Y [][]float64, c chan [][]float64) {
	var row []float64
	var prod [][]float64
	l1 := len(X)
	l2 := len(Y)
	if len(X[0]) != l2 {
		panic("Matrix dimension error!")
	}
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			row = append(row, VectorProduct(X[i], Y[j]))
		}
		prod = append(prod, row)
	}
	c <- prod
}

func ParallelMatrixProd(X, Y [][]float64) (u [][]float64) {
	l1 := len(X)
	l2 := len(Y)
	if len(X[0]) != l2 {
		panic("Matrix dimension error!")
	}
	n := runtime.NumCPU() * 2
	N := int(math.Floor(float64(l1 / n)))
	c := make(chan [][]float64)
	for i := 0; i < n; i++ {
		j1 := i*N + 1
		j2 := (i + 1) * N
		go MatrixProduct(X[j1:j2], Y, c)
	}

	if n*N != len(X) {
		go tMatrixProduct(X[(n*N):(len(X)-1)], Y, c)
	}

	for i := 0; i < n; i++ {
		submatrix := <-c
		u = append(u, submatrix...)
	}
	return u
}
