package main

import "fmt"

func main() {
  // Define map
  emails := make(map[string]string)

  // Assign key values
  emails["Dennis"] = "dennis@gmail.com"
  emails["Brad"] = "brad@gmail.com"
  emails["Adam"] = "adam@gmail.com"

  // Declare map and add key values
  emails := map[string]string{"Brad":"brad@gmail.com", "Dennis":"dennis@gmail.com"}

  fmt.Println(emails)
  // length of the map and how many items are in it:
  fmt.Println(len(emails))
  fmt.Println(emails["Dennis"])

  // Delete from map
  delete(emails, "Brad")
  fmtPrintln(emails)
}
