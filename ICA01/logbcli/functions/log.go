package functions

import (
	"math"
	"fmt"
	"os"
	"strconv"
)

// Regner ut log2 av paramteret
func Log2 (x float64) float64 {
	var result = math.Log2(x)
	return result
}

// regner ut log10 av parameteret
func Log10 (x float64) float64 {
	var result = math.Log10(x)
	return result
}
// tar to argumenter fra kommandolinjen. Arg brukes til å regne ut hvilken logbase vi skal bruke
// det andre argumentet brukes som parameter i regnestykket.
func CalcLog () {

	arg := os.Args[1]
	arg1 := os.Args[2]

	param, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		fmt.Println(err)
	}

	if arg == "log2" {
		fmt.Println(Log2(param))
	} else if arg == "log10" {
		fmt.Println(Log10(param))
	} else {
		fmt.Println("Invalid input. Please check if you entered valid arguments.")
		fmt.Printf(arg)
		fmt.Println(param)
	}
}