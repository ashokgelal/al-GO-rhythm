package main

import "fmt"

func main(){
  a, b := order(7, 2)
  fmt.Printf("The correct order: %v, %v\n", a, b)
}

func order(a, b int) (int, int){
  if a > b{
    return b,a
  }
  return a, b
}
