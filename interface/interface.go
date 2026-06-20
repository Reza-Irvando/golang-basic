package main

import ("math"
"fmt")

type hitung interface {
	keliling() float64
	luas() float64
}

type persegi struct {
	sisi float64
}

type lingkaran struct{
	jarijari float64
}

func (p persegi) keliling() float64 {
	return 4*p.sisi
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (l lingkaran) keliling() float64 {
	return math.Phi*l.jarijari*2
}

func (l lingkaran) luas() float64 {
	return math.Phi * math.Pow(l.jarijari, 2)
}

func main(){
	var bangunDatar hitung
	bangunDatar = persegi {10}
	fmt.Println("Persegi")
	fmt.Println("Keliling: ", bangunDatar.keliling())
	fmt.Println("Luas: ", bangunDatar.luas())

	bangunDatar = lingkaran {10}
	fmt.Println("Lingkaran")
	fmt.Println("Keliling: ", bangunDatar.keliling())
	fmt.Println("Luas: ", bangunDatar.luas())
}