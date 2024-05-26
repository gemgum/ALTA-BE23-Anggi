package main

import "fmt"

func CompareString(kata1 string, kata2 string) string {
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

	fmt.Println(CompareString("AKA", "AKASHI"))
	fmt.Println(CompareString("KANGOORO", "KANG"))
	fmt.Println(CompareString("KI", "KIJANG"))
	fmt.Println(CompareString("KUPU-KUPU", "KUPU"))
	fmt.Println(CompareString("ILALANG", "ILA"))

}
