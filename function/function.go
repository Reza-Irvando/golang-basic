package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Alice", "Wonderland"}
	sendMessage("Hello, ", names)

	// multiple return values
	a, b := persegi(5,10)
	fmt.Println("Keliling:", a)
	fmt.Println("Luas:", b)

	// variadic function
	avg := average(1, 2, 3, 4, 5)
	fmt.Println("Average of numbers:", avg)

	// closure function
	var hasilBagi = func(arrBagi [2]int) (int, int) {
		x := arrBagi[0]
		y := arrBagi[1]
		habisdibagi := x / y
		sisabagi := x % y
		return habisdibagi, sisabagi
	}

	pembagian := [2]int{10, 3}
	jumlah, sisa := hasilBagi(pembagian)
	fmt.Println("Hasil bagi:", jumlah)
	fmt.Println("Sisa bagi:", sisa)

	// callback function
	var hasil = filter([]string{"ini", "data"}, func(each string) bool {
		return true
	})
	fmt.Println("Hasil filter:", hasil)
}

func sendMessage(message string, arr []string) {
	// Function implementation goes here
	var fullName = strings.Join(arr, " ")
	fmt.Println(message, fullName) 
}

func persegi(panjang int, lebar int) (int, int) {
	keliling := 2 * (panjang + lebar)
	luas := panjang * lebar
	return keliling, luas
}

func average(nums ...int) float64 {
	var total int = 0
	for _, num := range nums {
		total += num
	}
	var avg float64 = float64(total) / float64(len(nums))
	return avg
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}