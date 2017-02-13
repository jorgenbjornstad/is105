package main

import "os"
import "fmt"
import "is105/ICA02/sum"

func main() {
  arg1 := os.Args[1]
  arg2 := os.Args[2]
  arg3 := os.Args[3]

if (arg1 == "int32"){
  var1 := sum.ConvertInt(arg2)
  fmt.Println("+")
  var2 := sum.ConvertInt(arg3)
  tall1 := int32(var1)
  tall2 := int32(var2)
  resultat :=  sum.SumInt32(tall1, tall2)
  fmt.Println("= ", resultat)
  }else if (arg1 == "int64"){
    tall1 := sum.ConvertInt(arg2)
    fmt.Println("+")
    tall2 := sum.ConvertInt(arg3)
    resultat :=  sum.SumInt64(tall1, tall2)
    fmt.Println("= ", resultat)
  }else if (arg1 == "uint32"){
    var1 := sum.ConvertUint(arg2)
    fmt.Println("+")
    var2 := sum.ConvertUint(arg3)
    tall1 := uint32(var1)
    tall2 := uint32(var2)
    resultat :=  sum.SumUint32(tall1, tall2)
    fmt.Println("= ", resultat)
  }else if (arg1 == "uint64"){
    tall1 := sum.ConvertUint(arg2)
    fmt.Println("+")
    tall2 := sum.ConvertUint(arg3)
    resultat :=  sum.SumUint64(tall1, tall2)
    fmt.Println("= ", resultat)
    }else {
    fmt.Println("Invalid type. Valid types are int32, int64, uint32, uint64")
    }
  }
