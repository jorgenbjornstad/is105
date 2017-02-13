package main
// Når denne pakken heter noe annet enn main, skjer det ingenting når
// "go build oppg3.int64.go" skrives i kommandolinjen. Altså det lages ikke en .exe fil.

import "os"
import "fmt"
import "strconv"
import "gitClone/is105/ICA02/sum"

func main() {

arg1 := os.Args[1]
arg2 := os.Args[2]


f1, err := strconv.ParseInt(arg1, 10, 64)
if err == nil {
  fmt.Println("  ",f1)
}else{
  fmt.Println("error")
}


f2, err2 := strconv.ParseInt(arg2, 10, 64)
if err2 == nil {
  fmt.Println("+ ",f2)
}else{
  fmt.Println("error")
}

resultat :=  sum.SumInt64(f1, f2)
fmt.Println("= ", resultat)

}
