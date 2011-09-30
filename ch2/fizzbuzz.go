package main

import "fmt"

func main(){
  const (
    FIZZ=3
    BUZZ=5
  )

  var isFizzBuzz bool

  for i:=1; i<100; i++{
    isFizzBuzz = false 
    if(i%FIZZ == 0){
      fmt.Printf("FIZZ")
      isFizzBuzz = true
    }

    if(i%BUZZ == 0){
      fmt.Printf("BUZZ")
      isFizzBuzz = true
    }

    if(!isFizzBuzz){
      fmt.Printf("%v", i)
    }
  }
}
