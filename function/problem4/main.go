package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func encryptData(kata string) {

	var hasilAkhir string
	for i := 0; i < len(kata); i++ {

		if kata[i] >= 65 && kata[i] <= 90 {
			if kata[i] >= 81 {
				hasilAkhir = hasilAkhir + string(kata[i]-16)
			} else {
				hasilAkhir = hasilAkhir + string(kata[i]+10)
			}

		} else {
			hasilAkhir = hasilAkhir + string(kata[i])

		}

	}

	fmt.Println(hasilAkhir)

}

func main() {
	// Problem 4 - Ubah Huruf
	fmt.Print("\nUbah Huruf\n")
	fmt.Printf("Masukan Kata = ")

	reader := bufio.NewReader(os.Stdin)
	kata, _ := reader.ReadString('\n')
	kata = strings.TrimSpace(kata)

	encryptData(kata)

}
