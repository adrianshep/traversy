package main

import "fmt"

func main() {
  var age int = 18
  var isCool = true
  isCool = false

  // name := "Adam"
  // email := "adam@gmail.com"

  name, email := "Adam", "adam@gmail.com"
  size := 1.3

  fmt.Println(name, age, isCool)
  fmt.Printf("%T\n", isCool)
}
