package math

import (
	"math"
	"runtime"

	"github.com/entropyx/tools"
)

// Mean function
func Mean(x []float64) float64 {
	out := 0.00
	n := len(x)
	for i := 0; i < n; i++ {
		out = out + x[i]
	}
	out = out / float64(n)
	return out
}

// Sd standart desviation function
func Sd(x []float64) float64 {
	mu := Mean(x)
	l := len(x)
	out := 0.00
	for i := 0; i < l; i++ {
		out = out + math.Pow(x[i]-mu, 2)
	}
	out = math.Sqrt(out)
	return out
}

// Min of vector
func Min(x []float64) float64 {
	min := x[0]
	l := len(x)
	for i := 0; i < l; i++ {
		if x[i] < min {
			min = x[i]
		}
	}
	return min
}

// Max of vector
func Max(x []float64) float64 {
	max := x[0]
	l := len(x)
	for i := 0; i < l; i++ {
		if x[i] > max {
			max = x[i]
		}
	}
	return max
}

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

// Normalize scale data.
func Normalize(X [][]float64) [][]float64 {
	l1 := len(X)
	l2 := len(X[0])
	mu := tools.Apply(X, 2, Mean)
	sigma := tools.Apply(X, 2, Sd)

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			X[i][j] = (X[i][j] - mu[j]) / sigma[j]
		}
	}
	return X
}

// MinMax scale data.
func MinMax(X [][]float64) [][]float64 {
	l1 := len(X)
	l2 := len(X[0])
	min := tools.Apply(X, 2, Min)
	max := tools.Apply(X, 2, Max)

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			X[i][j] = (X[i][j] - min[j]) / (max[j] - min[j])
		}
	}
	return X
}

// RescaleCoef rescale coefficients fitted.
func RescaleCoef(theta, mu, sigma []float64) []float64 {
	l := len(theta)
	for i := 0; i < l-1; i++ {
		theta[0] = theta[0] - theta[i+1]*mu[i]/sigma[i]
		theta[i+1] = theta[i+1] / sigma[i]
	}
	return theta
}

// Round array with some presicion
func Round(num []float64, precision int) []float64 {
	var out []float64
	v1 := math.Pow(10, float64(precision))
	for i := 0; i < len(num); i++ {
		v2 := int(num[i]*v1 + math.Copysign(0.5, num[i]*v1))
		out = append(out, float64(v2)/v1)
	}
	return out
}
