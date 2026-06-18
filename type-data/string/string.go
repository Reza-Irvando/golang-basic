package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "Reza"
	deskripsi := "Reza adalah seorang programmer yang sedang belajar bahasa pemrograman Go."

	

	// Print the values
	fmt.Println("Nama:", name)
	fmt.Println("Deskripsi:", deskripsi)

	// String Operations
	text := "Hello, World"
	fmt.Println("Text:", text)

	fmt.Println("Length of text:", len(text))
	fmt.Println("Uppercase text:", strings.ToUpper(text))
	fmt.Println("Lowercase text:", strings.ToLower(text))
	fmt.Println("Starts with 'Hello':", strings.HasPrefix(text, "Hello"))
	fmt.Println("Starts with 'World':", strings.HasPrefix(text, "World"))
	fmt.Println("Contains 'World' in text:", strings.Contains(text, "World"))
	fmt.Println("Contains 'Go' in text:", strings.Contains(text, "Go"))
	newText := strings.ReplaceAll(text, "World", "Go")
	fmt.Println("Replace 'World' with 'Go':", newText)

	items := strings.Split("apple,banana,cherry", ", ")
	fmt.Println("Items:", items)
}