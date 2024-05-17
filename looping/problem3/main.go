package main

import "fmt"

func main() {
	//Problem 3 - Bilangan Prima
	var angka int
	var isPrima = true
	fmt.Print("\nCek Bilangan Prima\n\n")
	fmt.Printf("Masukan Angka = ")
	fmt.Scanln(&angka)

	for i := 2; i < angka; i++ {
		if angka%i == 0 {
			isPrima = false
			break
		}
	}

	if angka == 1 {
		isPrima = false
	}
	if isPrima {
		fmt.Println("Angka", angka, "Merupakan Bilanga Prima")
	} else {
		fmt.Println("Angka", angka, "Merupakan Bukan Bilanga Prima")
	}
}
