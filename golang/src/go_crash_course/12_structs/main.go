package main

import ("fmt"
        "strconv")

// Define person struct
type Person struct {
  // firstName string
  // lastName string
  // city string
  // gender string
  // age int

  firstName, lastName, city, gender string
  age                               int
}

// Greeting method (value receiver)
// have to convert int age to string for it this to work
func (p Person) greet() string {
  return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
  p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
  if p.gender == "m" {
    return
  } else {
    p.lastName = spouseLastName
  }
}

func main() {
  // initialize person using struct
  person1 := Person{firstName: "Emily", lastName: "Pleiad", city: "New Meeting", gender: "f", age: 18}

  // alternative
  person2 := Person{"Adam","Jura", "Rogerstowne", "m", 18}

  fmt.Println(person1.firstName)
  person1.age++
  fmt.Println(person1)

  person1.hasBirthday()
  person1.hasBirthday()
  person1.getMarried("Jura")

  person2.getMarried("Pleiad")

  fmt.Println(person1.greet())
}
