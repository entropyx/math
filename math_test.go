package math

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMathFunction(t *testing.T) {
	Convey("Given the following dataset ...", t, func() {
		var data [][]float64
		var X [][]float64
		filePath := "/home/gibran/Work/Go/src/github.com/entropyx/math/dataset/dataset2.txt"
		strInfo, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		trainingData := strings.Split(string(strInfo), "\n")
		for _, line := range trainingData {
			if line == "" {
				break
			}

			var values []float64
			for _, value := range strings.Split(line, " ") {
				floatVal, err := strconv.ParseFloat(value, 64)
				if err != nil {
					panic(err)
				}
				values = append(values, floatVal)
			}
			data = append(data, values)
		}

		for i := 0; i < len(data); i++ {
			X = append(X, data[i][:3])
		}

		Convey("The Traspose of matrix X ... ", func() {
			t := Traspose(X)
			So(len(t[0]), ShouldEqual, len(X))
		})

		Theta := [][]float64{make([]float64, len(X[0]))}
		c := make(chan [][]float64)
		go MatrixProduct(X, Theta, c)
		prod := <-c
		prod2 := ParallelMatrixProd(X, Theta)

		var out1 []float64
		var out2 []float64
		for i := 0; i < len(X); i++ {
			out1 = append(out1, prod[i][0])
			out2 = append(out2, prod2[i][0])
		}

		Convey("The product between matrix X an Theta ... ", func() {
			So(out1, ShouldResemble, make([]float64, len(X)))
		})

		Convey("The parallel product between matrix X an Theta ... ", func() {
			So(out2, ShouldResemble, make([]float64, len(X)))
		})

		coef1 := []float64{-25.1613335, 0.2062317, 0.2014716}
		coef2 := []float64{1.718449, 4.012903, 3.743903}
		mu := []float64{65.64427, 66.22200}
		sigma := []float64{19.45822, 18.58278}

		Convey("The rescaling of coef2 shoul be equal to coef1 ", func() {
			rescale := RescaleCoef(coef2, mu, sigma)
			So(Round(rescale, 4), ShouldResemble, Round(coef1, 4))
		})
	})
}
