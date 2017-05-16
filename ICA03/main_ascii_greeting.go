package main

import "./ascii"
import "fmt"

func main() {
print1, print2, print3 := ascii.GreetingASCII()
fmt.Printf("Den binære representasjonen av \"hello\": %b \n", print1, )
fmt.Printf("Den hexadesimale representasjonen av \"hello\": %x \n", print2)
fmt.Printf("%s \n", print3)

}
