package main

import "fmt"

func CaesarCipher(off int, kata string) string {
	var hasilAkhir string

	for _, v := range kata {
		var kata_pos int = int(v)
		for j := 0; j < off; j++ {
			kata_pos++
			if kata_pos > 122 {
				kata_pos = 97
			}
		}
		hasilAkhir += fmt.Sprint(string(kata_pos))
	}
	return hasilAkhir
}

func main() {
	// Problem 2 - Caesar Cipher
	fmt.Println(CaesarCipher(3, "abc"))
	fmt.Println(CaesarCipher(2, "alta"))
	fmt.Println(CaesarCipher(10, "alterraacademy"))
	fmt.Println(CaesarCipher(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(CaesarCipher(1000, "abcdefghijklmnopqrstuvwxyz"))
}
