package main

import "fmt"

func main(){
  str := "foobar"
  byteArr := []byte(str)

  for i,j := 0, len(byteArr)-1; i<j; i,j = i+1, j-1{
    byteArr[i], byteArr[j] = byteArr[j], byteArr[i]
  }

  fmt.Printf("%s\n", string(byteArr))
}
