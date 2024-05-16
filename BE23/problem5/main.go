package main

import "fmt"

func main() {
	//Problem 5 - Menghitung Luas Permukaan Tabung
	var T float64 = 20
	var r float64 = 4

	const pi = 3.14

	fmt.Print("\nPenghitung Luas Permukaan Tabung\n\n")
	fmt.Printf("Masukan Tinggi Tabung = ")
	fmt.Scanln(&T)

	fmt.Printf("Masukan Radius Tabung = ")
	fmt.Scanln(&r)

	hasil := 2*pi*r*r + 2*pi*r*T
	fmt.Printf("\nMaka Luas Permukaan Tabung = %.2f\n\n", hasil)

}
