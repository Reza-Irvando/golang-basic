package main

import "fmt"

func main() {
	// Operasi Perbandingan
	var a, b int
	fmt.Print("Input a: ")
	fmt.Scanln(&a)
	fmt.Print("Input b: ")
	fmt.Scanln(&b)

	fmt.Println("Hasil Perbandingan:")
	fmt.Printf("%d == %d ? %v\n", a, b, a == b)
	fmt.Printf("%d != %d ? %v\n", a, b, a != b)
	fmt.Printf("%d > %d ? %v\n", a, b, a > b)
	fmt.Printf("%d < %d ? %v\n", a, b, a < b)
	fmt.Printf("%d >= %d ? %v\n", a, b, a >= b)
	fmt.Printf("%d <= %d ? %v\n", a, b, a <= b)

	// Operasi Logika
	fmt.Println("\nHasil Operasi Logika:")
	fmt.Printf("5 > 7 || 9 < 7: %v\n", 5 > 7 || 9 < 7)
	fmt.Printf("5 < 7 || 9 > 7: %v\n", 5 < 7 || 9 > 7)
	fmt.Printf("5 > 7 && 9 < 7: %v\n", 5 > 7 && 9 < 7)
	fmt.Printf("5 < 7 && 9 > 7: %v\n", 5 < 7 && 9 > 7)
	fmt.Printf("!(5 > 7): %v\n", !(5 > 7))
}