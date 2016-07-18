package main

import (
  "fmt"
  "math"
)

func main() {

  userInput := 20  

  binDigits := [10]float64

  binDigits = binaryConv(userInput, binDigits)

  fmt.Println("Your number converted to binary is %r\n", binDigits)
  
  }

  func squareOf(input float64) float64 {
    return Sqrt(input)
  }

  func byTwo(input float64) float64 {
    return input / 2
  }

  func rem(input float64) float64 {
    return input % 2
  }

  func binaryConv(entry float64, slicefloat64 []float64) []float64 {
    var i = 0
    var leftOver = byTwo
    for leftOver >= 0 {
      for i < entry {
          powTwoOfEntry := squareOf(entry)
          slicefloat64[powTwoOfEntry] = 1
          leftOver = byTwo - entry^2 //decimal number
          rem := rem(leftOver)
          if leftOver == 0 {
            if rem == 0 {
            break
            } else {
              slicefloat64[0] = 1
              break
            }
            
          }
          entry = byTwo(leftOver)
          }
        return slicefloat64
      }
    }
