package main

import "fmt"

func main() {
	//Problem 2.2 - Faktor Bilangan
	var angka int
	fmt.Print("\nPencari Faktor Bilangan dari Terbesar \n\n")
	fmt.Printf("Masukan Angka = ")
	fmt.Scanln(&angka)

	for i := angka; i > 0; i-- {
		if angka%i == 0 {
			println(i)

		}
	}

}
