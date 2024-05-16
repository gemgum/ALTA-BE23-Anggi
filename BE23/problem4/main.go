package main

import "fmt"

func main() {
	//Problem 4 - Menghitung Luas Segitiga
	var alas float64 = 20
	var tinggi float64 = 25

	fmt.Print("\nPenghitung Luas Segitiga\n\n")
	fmt.Printf("Masukan Alas Segitiga = ")
	fmt.Scanln(&alas)

	fmt.Printf("Masukan Tinggi Segitiga = ")
	fmt.Scanln(&tinggi)

	hasil := alas * tinggi * 0.5
	fmt.Printf("\nMaka Luas Segitiga %.0f\n\n", hasil)

}
