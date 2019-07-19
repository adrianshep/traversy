package main

import ("fmt"
        "strconv")

// Define person struct
type Person struct {
  firstName string
  lastName string
  city string
  gender string
  age int
}

// Greeting method (value receiver)
// have to convert int age to string for it this to work
func (p Person) greet() string {
  return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func main() {
  // initialize person using struct
  person1 := Person{firstName: "Emily", lastName: "Pleiad", city: "New Meeting", gender: "f", age: 18}

  // alternative
  // person1 := Person{"Emily","Pleiad", "New Meeting", "f", 18}

  fmt.Println(person1.firstName)
  person1.age++
  fmt.Println(person1)

  fmt.Println(person1.greet())
}
