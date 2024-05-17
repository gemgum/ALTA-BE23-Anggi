package main

import "fmt"

func main() {
	// Problem 1 - Konversi Nilai

	var nilai int8
	fmt.Print("\nConversi Nilai\n\n")
	fmt.Printf("Masukan Nilai = ")
	fmt.Scanln(&nilai)

	var nilaiChar string
	if nilai <= 100 && nilai >= 80 {
		nilaiChar = "A"
	} else if nilai <= 79 && nilai >= 65 {
		nilaiChar = "B+"
	} else if nilai <= 64 && nilai >= 50 {
		nilaiChar = "B"
	} else if nilai <= 49 && nilai >= 35 {
		nilaiChar = "C"
	} else if nilai <= 34 && nilai >= 0 {
		nilaiChar = "D"
	} else {
		fmt.Println("Format Nilai yang Dimasukan Tidak Sesuai ")
	}

	fmt.Println("Nilai yang didapatkan adalah ", nilaiChar)
}
