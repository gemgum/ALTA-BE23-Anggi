package main

import "fmt"

func main() {
	//Problem 5 - Exponentiation
	var x, n int
	fmt.Print("\nHitung Pangkat\n\n")
	fmt.Printf("Masukan Angka = ")
	fmt.Scanln(&x)
	fmt.Printf("Masukan Pangkat = ")
	fmt.Scanln(&n)

	var hasil int = x
	for i := 1; i < n; i++ {
		hasil = hasil * x

	}

	fmt.Println(x, "pangkat", n, "adalah", hasil)

}
