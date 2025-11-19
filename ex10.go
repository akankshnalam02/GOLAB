package main
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/go-gota/gota/dataframe"
	"github.com/sajari/regression"
)
func main() {
	passengersFile, err := os.Open("log_diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()
	passengersDF := dataframe.ReadCSV(passengersFile)
	passengers := passengersDF.Col("log_differenced_passengers").Float()
	coeffs, intercept := autoregressive(passengers, 2)
	fmt.Printf("\nlog(x(t))- log(x(t-1)) = %0.6f + lag1*%0.6f + lag2*%0.6f\n\n", intercept, coeffs[0], coeffs[1])
}
func autoregressive(x []float64, lag int) ([]float64, float64) {
	var r regression.Regression
	r.SetObserved("x")
	for i := 0; i < lag; i++ {
		r.SetVar(i, "x"+strconv.Itoa(i))
	}
	xAdj := x[lag:len(x)]
	for i, xVal := range xAdj {
		laggedVariables := make([]float64, lag)
		for idx := 1; idx <= lag; idx++ {
			laggedVariables[idx-1] = x[lag+i-idx]
		}
		r.Train(regression.DataPoint(xVal, laggedVariables))
	}
	r.Run()
	var coeff []float64
	for i := 1; i <= lag; i++ {
		coeff = append(coeff, r.Coeff(i))
	}
	return coeff, r.Coeff(0)
}
