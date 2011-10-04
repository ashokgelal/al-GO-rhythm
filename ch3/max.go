package main

import "fmt"

func max(l []int) (max int){
  max = l[0]

  for _, v := range l{
    if v > max{
      max = v
    }
  }

  return
}

func main(){
  data := []int {4, 18, 2, 1, 5, 6 }
  fmt.Printf("Max: %v\n", max(data))
}
