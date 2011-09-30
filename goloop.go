package main

import "fmt"

func main(){
  var i=0
  Lbl:{
    fmt.Printf("%d\n", i)
    i++
    if(i<10){
      goto Lbl
    }
  }
}
