package main

import (
	"fmt"
	"time"
)

func prosesCepat(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Data dari proses CEPAT"
}

func prosesLambat(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "Data dari proses LAMBAT"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go prosesCepat(ch1)
	go prosesLambat(ch2)

	fmt.Println("Menunggu data masuk...")

	// select akan menunggu sampai ch1 ATAU ch2 mengirim data
	for i:=0; i<3; i++{
		select {
		case hasil1 := <-ch1:
			fmt.Println("Diterima:", hasil1)
		case hasil2 := <-ch2:
			fmt.Println("Diterima:", hasil2)
		case <-time.After(2 * time.Second):
			fmt.Println("Error: Waktu habis (Timeout)! API terlalu lambat.")
		}
	}
	
	// Output: 
	// Menunggu data masuk...
	// Diterima: Data dari proses CEPAT
}