package main

import (
  "fmt"
  "utf8"
)

func main(){
  var str = "asSASA ddd dsjkdsjs dk"
  fmt.Printf("String %s\nLength: %d, Runes: %d\n", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
}
