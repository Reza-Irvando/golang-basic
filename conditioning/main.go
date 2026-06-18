package main

import "fmt"

func main() {
	// If else statement
	var height int
	fmt.Printf("Masukkan tinggi badan (dalam cm): ")
	fmt.Scanln(&height)
	if height > 175 {
		fmt.Println("Anda termasuk dalam kategori tinggi.")
	} else if height >= 160 {
		fmt.Println("Anda termasuk dalam kategori standar.")
	} else {
		fmt.Println("Anda termasuk dalam kategori pendek.")
	}

	// Switch statement
	var day int
	fmt.Printf("Masukkan nomor hari (1-7): ")
	fmt.Scanln(&day)
	switch day {
	case 1:
		fmt.Println("Hari Senin")
	case 2:
		fmt.Println("Hari Selasa")
	case 3:
		fmt.Println("Hari Rabu")
	case 4:
		fmt.Println("Hari Kamis")
	case 5:
		fmt.Println("Hari Jumat")
	case 6:
		fmt.Println("Hari Sabtu")
	case 7:
		fmt.Println("Hari Minggu")
	default:
		fmt.Println("Nomor hari tidak valid")
	}
}