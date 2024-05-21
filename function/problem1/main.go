package main

import "fmt"

func playWithAsterisk(angka int) {
	var spasi = angka
	var bintang int
	for i := 0; i <= angka; i++ {
		for j := 0; j < spasi; j++ {
			fmt.Print(" ")
		}
		bintang++
		for k := 0; k < bintang; k++ {
			fmt.Print("* ")
		}

		fmt.Println("")
		spasi--

	}
}

func main() {
	// Problem 1 - Play With Asterisk
	var angka int
	fmt.Print("\nPlay With Asterisk \n\n")
	fmt.Printf("Masukan Angka = ")
	fmt.Scanln(&angka)

	playWithAsterisk(angka)

}
