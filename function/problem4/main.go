package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Problem 4 - Kata Palindrome
	fmt.Print("\nCek Kata Palindrome\n")
	fmt.Printf("Masukan Kata = ")

	reader := bufio.NewReader(os.Stdin) // masukan keyboard
	kata, _ := reader.ReadString('\n')  // baca string yang tersedia
	kata = strings.TrimSpace(kata)      // hilangkan '\n'

	var kataBantu string

	// Reverse kata pindah ke kataBantu
	for _, v := range kata {
		kataBantu = string(v) + kataBantu
	}

	var palindrome int = strings.Compare(kata, kataBantu)
	if palindrome != 0 {
		fmt.Println("Kata", kata, "Merupakan Bukan Kata Palindrome")
	} else {
		fmt.Println("Kata", kata, "Merupakan Kata Palindrome")

	}

}
