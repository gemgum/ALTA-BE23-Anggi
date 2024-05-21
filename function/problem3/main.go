package main

import "fmt"

func tabelPerkalian(angka int) {
	var spasi = angka
	var hasil int
	for i := 1; i < angka; i++ {
		for j := 1; j < spasi; j++ {
			hasil = i * j

			fmt.Print(hasil, "\t| ")

		}
		fmt.Println("")

	}
}

func main() {
	// Problem 3 - Cetak Tabel Perkalian
	var angka int
	fmt.Print("\nCetak Tabel Perkalian\n\n")
	fmt.Printf("Masukan Angka = ")
	fmt.Scanln(&angka)

	tabelPerkalian(angka)

}
