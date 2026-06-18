package main

import "fmt"

func main() {
	// Operasi Matematika
	var a int = 10
	var b int = 5
	fmt.Println("Penjumlahan:", a+b)
	fmt.Println("Pengurangan:", a-b)
	fmt.Println("Perkalian:", a*b)
	fmt.Println("Pembagian:", a/b)
	fmt.Println("Sisa Bagi:", a%b)

	// Assignment Operators
	a += 2
	fmt.Println("a += 2 -> ", a)
	a -= 4
	fmt.Println("a -= 4 -> ", a)
	a *= 2
	fmt.Println("a *= 2 -> ", a)
	a /= 4
	fmt.Println("a /= 4 -> ", a)
	a %= 3
	fmt.Println("a %= 3 -> ", a)

	// Increment & Decrement
	a++
	fmt.Println("a++ -> ", a)
	a--
	fmt.Println("a-- -> ", a)
}