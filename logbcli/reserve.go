package logbcli

import (
	"flag"
	"fmt"
)


logbcli
func main() {
logBase := flag.Float64("Base", 0, "hvilken log base du vil bruke")
param := flag.Float64("Parameter", 0, "tallet du vil regne ut logaritmen av")
flag.Parse()


var log2 float64 = 2
var log10 float64 = 10

if *logBase == log2 {
fmt.Println(functions.Log2(*param))
} else if *logBase == log10 {
fmt.Println(functions.Log10(*param))
} else  {
fmt.Println("Not valid")
}

}