package main

import (
	"errors"
	"fmt"
	"strconv"
)

func validate(genap int) (bool, error) {
	hasil := genap%2
	if hasil != 0 {
		return false, errors.New("input bukan angka genap")
	}
	return true, nil
}

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
		fmt.Println(input, "bukan angka yang valid.")
		fmt.Println("Terjadi kesalahan:", err)
	}

	var input2 int
	fmt.Println("Masukkan angka genap:")
	fmt.Scanln(&input2)
	if valid, err := validate(input2); valid {
		fmt.Println(input2, "adalah angka genap.")
	} else {
		fmt.Println(err)
	}
}