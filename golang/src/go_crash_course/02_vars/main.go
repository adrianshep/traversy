package main

import "fmt"

func main() {
  var age int = 18
  var isCool = true
  isCool = false

  name := "Adam"
  size := 1.3

  fmt.Println(name, age, isCool)
  fmt.Printf("%T\n", isCool)
}
