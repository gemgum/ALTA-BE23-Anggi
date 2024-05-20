package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Problem 6 - Full Prima
	var angkaS string
	var isPrima, isFullPrima = true, true

	fmt.Print("\nFull Bilangan Prima\n\n")
	fmt.Printf("Masukan Angka = ")
	fmt.Scanln(&angkaS)

	// Konversi dari string menjadi int
	angka, err := strconv.Atoi(angkaS)

	if err != nil {
		panic(err)
	}

	// Flow pengecekan bilangan prima
	for i := 2; i < angka; i++ {
		if angka%i == 0 {
			isFullPrima, isPrima = false, false
			break
		}
	}

	if angka == 1 {
		isPrima = false
	}

	if isPrima {
		for i := 0; i < len(angkaS); i++ {

			var n string = string(angkaS[i])

			// Konversi dari string menjadi int
			m, erro := strconv.Atoi(n)

			if erro != nil {
				panic(erro)
			}

			// Cek apakah salah satu digit sama dengan 1
			if m == 1 {
				isFullPrima = false
				break
			}

			// Flow pengecekan bilangan prima
			for j := 2; j < m; j++ {
				if m%j == 0 {
					isFullPrima = false
					break
				}
			}

			if !isFullPrima {
				break
			}
		}

	}

	if isFullPrima {
		fmt.Println("Angka", angka, "Merupakan Bilanga Full Prima")
	} else {
		fmt.Println("Angka", angka, "Merupakan Bukan Bilanga Full Prima")
	}

}
