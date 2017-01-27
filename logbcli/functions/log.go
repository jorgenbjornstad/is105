package functions

import (
	"math"
	"flag"
	"fmt"
)


func Log2 (x float64) float64 {
	var result = math.Log2(x)
	return result

}
func Log10 (x float64) float64 {
	var result = math.Log10(x)
	return result
}

func CalcLog () {

	logBase := flag.Float64("Base", 0, "hvilken log base du vil bruke")
	param := flag.Float64("Parameter", 0, "tallet du vil regne ut logaritmen av")
	flag.Parse()

	var log2 float64 = 2
	var log10 float64 = 10

	if *logBase == log2 {
		fmt.Println(Log2(*param))
	} else if *logBase == log10 {
		fmt.Println(Log10(*param))
	} else {
		fmt.Println("Not valid")
	}
}