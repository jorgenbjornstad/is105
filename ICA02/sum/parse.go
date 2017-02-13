package sum

import "strconv"
import "fmt"

func ConvertInt(input1 string) int64{
  value, err := strconv.ParseInt(input1, 10, 64)
  if err == nil {
    fmt.Println(value)
  }else{
    fmt.Println("error")
  }
  return value
  }

func ConvertUint(input2 string) uint64{
    value, err := strconv.ParseUint(input2, 10, 64)
    if err == nil {
      fmt.Println(value)
    }else{
      fmt.Println("error")
    }
    return value
    }
func ConvertFloat(input3 string) float64{
    value, err := strconv.ParseFloat(input3, 10, 64)
    if err == nil {
        fmt.Println(value)
     }else{
          fmt.Println("error")
      }
      return value
      }
