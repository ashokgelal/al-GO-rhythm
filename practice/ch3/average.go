package main

import "fmt"

func main(){
  data := []float64{2.5, 2.5, 5}
  avg := average(data)
  fmt.Printf("The average is: %v", avg)
}

func average(xs []float64)(ave float64){
  sum := 0.0
  switch len(xs){
    case 0:
      ave = 0
    default:
      for _, v := range xs {
        sum+= v
      }
      ave = sum/float64(len(xs))
  }
  return
}
