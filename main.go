package main

import "fmt"

func print_age(age int) {
  fmt.Printf("Age: %v\n", age)
}

func main() {
  age := 36
  print_age(age)
}
