package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Println("Masukkan angka:")
	fmt.Scanln(&input)
	var angka int
	var err error
	angka, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(angka, "adalah angka yang valid.")
	} else {
		defer fmt.Println(input, "bukan angka yang valid.")
		panic(err.Error())
	}
}