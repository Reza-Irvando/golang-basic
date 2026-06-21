package main

import "fmt"

func main() {
	// Membuat channel integer dengan kapasitas (buffer) sebanyak 3
	ch := make(chan int, 3)

	// Pengirim tidak akan ter-block karena masih ada ruang (buffer)
	ch <- 10
	ch <- 20
	ch <- 30
	
	ch <- 40 // ❌ JIKA INI DIBUKA, PROGRAM AKAN ERROR (DEADLOCK) KARENA BUFFER PENUH

	data1 := <- ch
	data2 := <- ch
	data3 := <- ch

	// Penerima mengambil data sesuai urutan masuk (First In, First Out)
	fmt.Println(data1) // Output: 10
	fmt.Println(data2) // Output: 20
	fmt.Println(data3) // Output: 30
}