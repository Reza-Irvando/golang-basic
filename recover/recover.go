package main

import (
	"fmt"
	"strconv"
)

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Terjadi kesalahan:", r)
	}	else {
		fmt.Println("Program berjalan dengan baik.")
	}
}

func main() {
	defer catch()
	var input string
	fmt.Println("Masukkan angka:")
	fmt.Scanln(&input)
	var angka int
	var err error
	angka, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(angka, "adalah angka yang valid.")
	} else {
		panic(err.Error())
	}
}