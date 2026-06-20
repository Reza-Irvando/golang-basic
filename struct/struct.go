package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Bob", Age: 25}
	fmt.Printf("%s is %d years old.\n", p1.Name, p1.Age)
	fmt.Printf("%s is %d years old.\n", p2.Name, p2.Age)
}