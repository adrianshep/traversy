package main

import "fmt"

func main() {
  a := 5
  b := &a

  fmt.Println(a, b)
  // below results in int
  fmt.Printf("%T\n", a)
  // while below results in *int, a pointer, which is a different data type
  fmt.Printf("%T\n", b)

  // Use * to read value from address
  fmt.Println(*b)
  // memory address only, same as *&a
  fmt.Println(b)

  // Change value with pointer
  *b = 10
  fmt.Println(a)
}
