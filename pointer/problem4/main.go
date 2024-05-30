package main

import (
	"fmt"
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

type student struct {
	name       string
	nameEncode string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	// var nameEncode = ""
	for i := 0; i < len(s.name); i++ {

		if s.name[i]+3 > 122 {
			s.nameEncode = s.nameEncode + string(s.name[i]+3-26)
		} else {
			s.nameEncode = s.nameEncode + string(s.name[i]+3)
		}

	}
	// fmt.Println(s.nameEncode)
	return s.nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	for i := 0; i < len(s.nameEncode); i++ {
		if s.nameEncode[i]-3 < 97 {
			nameDecode = nameDecode + string(s.nameEncode[i]-3+26)
		} else {
			nameDecode = nameDecode + string(s.nameEncode[i]-3)
		}

	}
	// fmt.Println(nameDecode)
	return nameDecode
}

func main() {
	// Problem 4 - Substitution Cipher

	var menu int
	var a student
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
