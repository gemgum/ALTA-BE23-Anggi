package main

import "fmt"

func caesar(kata1 string, kata2 string) string {
	var value string
	if len(kata1) > len(kata2) {
		for i, _ := range kata2 {
			if kata2[i] == kata1[i] {
				value += string(kata1[i])
			}

		}
	} else {
		for i, _ := range kata1 {
			if kata1[i] == kata2[i] {
				value += string(kata2[i])
			}

		}
	}

	return value
}

func main() {
	// Problem 1 - Compare String

	fmt.Println(caesar("AKA", "AKASHI"))
	fmt.Println(caesar("KANGOORO", "KANG"))
	fmt.Println(caesar("KI", "KIJANG"))
	fmt.Println(caesar("KUPU-KUPU", "KUPU"))
	fmt.Println(caesar("ILALANG", "ILA"))

}
