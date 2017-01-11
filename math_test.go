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

		c1 := make(chan []float64)
		go MatrixVectorProduct(X, []float64{1, 2, 3}, c1)
		prod1 := <-c1
		close(c1)

		Theta := [][]float64{[]float64{1, 2, 3}}

		c2 := make(chan [][]float64)
		go MatrixProduct(X, Theta, c2)
		prod2 := <-c2
		close(c2)

		prod3 := ParallelMatrixProd(X, Theta)

		var out1 []float64
		var out2 []float64
		for i := 0; i < len(X); i++ {
			out1 = append(out1, prod2[i][0])
			out2 = append(out2, prod3[i][0])
		}

		result := []float64{304.3213976952, 193.2584141084, 291.4014116211, 380.29085505890004, 385.09860141210004, 260.1156702979, 412.7476067283, 290.71153175840004,
			415.4592825625, 300.46581985449995, 307.3989443158, 242.8371064821, 395.05999665520005, 432.8852534051, 308.187110841, 376.5641547097, 297.3616973162, 276.9294332743,
			421.10443279080005, 297.68545840440004, 264.2593700523, 377.75164928390006, 248.6370111626, 202.05270775190002, 363.7652628891, 335.4054012035, 295.84524829860004,
			303.6307886061, 275.4287357234, 273.5586504598, 342.1822408344, 342.9649907171, 294.59883077020004, 313.38938989810003, 294.9770315364, 266.91186875020003, 365.4393074448,
			372.07815950990005, 275.2987516034, 295.0804410608, 337.7289259733, 244.6643113287, 386.5935003552, 287.59227299329996, 240.563607916, 281.62832736760004, 366.7606698574,
			456.72666555, 415.45259996569996, 450.21882663270003, 383.3789939326, 382.5421468538, 312.2652336475, 251.2380550816, 250.98660867759997, 279.6002074161, 402.87598609279996,
			352.9500431779, 358.9710964142, 379.9536108528, 409.0710329513, 212.6337674722, 231.29205252600002, 209.896566495, 289.6867818864, 257.398083183, 374.5206584349,
			254.7947619003, 437.90733046959997, 317.467619221, 196.3671852741, 363.17370489079997, 434.3757673544, 341.2007657712, 346.25725737299996, 417.7634713952, 360.95613320380005,
			329.3458981952, 249.43634089250003, 293.6129587674, 388.2391595286, 327.7519338561, 335.40657447089995, 294.02025857169997, 434.61395569589996, 394.7199648853, 321.6852670815,
			423.2290210031, 448.2131302047, 288.0045237438, 420.6659775225, 444.4234772364, 218.6764326983, 404.5207953847, 316.7664642471, 313.1191828819, 346.83495444090005, 405.9562460265,
			306.4758491407, 419.14122468709996}

		Convey("The product between matrix X an vector [1 2 3] ... ", func() {
			So(prod1, ShouldResemble, result)
		})

		Convey("The product between matrix X an Matrix Theta ... ", func() {
			So(out1, ShouldResemble, result)
		})

		Convey("The parallel product between matrix X an Theta ... ", func() {
			So(out2, ShouldResemble, result)
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
