package main

import "fmt"

// Define person struct
type Person struct {
  firstName string
  lastName string
  city string
  gender string
  age int
}

func main() {
  // initialize person using struct
  person1 := Person{firstName: "Emily", lastName: "Pleiad", city: "New Meeting", gender: "f", age: 18}

  // alternative
  person1 := Person{"Emily","Pleiad", "New Meeting", "f", 18}

  fmt.Println(person1)
  fmt.Println(person1.firstName)
}
