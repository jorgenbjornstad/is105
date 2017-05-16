package functions

import "math"

//funksjon for å regne ut log2 av gitt tall i float64

func Log2(y float64) float64 {
	var x = y
	var result = math.Log2(x)
	return result
}
