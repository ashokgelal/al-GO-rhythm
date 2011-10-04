package main

import "fmt"

func main(){
  printAll(1, 2, 3, 5, 6, 4, 3)
  printAll(10, 11, 12)
}

func printAll(numbers ... int){
  for _, v := range numbers{
    fmt.Printf("%d\n", v)
  }
}
