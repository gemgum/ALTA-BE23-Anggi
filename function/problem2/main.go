package main

import "fmt"

func main() {
	//Problem 2.1 - Faktor Bilangan
	var angka int
	fmt.Print("\nPencari Faktor Bilangan dari Terkecil \n\n")
	fmt.Printf("Masukan Angka = ")
	fmt.Scanln(&angka)

	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			println(i)

		}
	}

}
