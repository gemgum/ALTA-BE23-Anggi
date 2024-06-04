package main

import (
	"fmt"
)

func isPrima(number int) bool {
	var isPrima bool = true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrima = false
			break
		}
	}

	if number == 1 {
		isPrima = false
	}
	return isPrima
}

func PrimetoX(number int) int {
	var pos int
	var isPos bool = false
	var tempNumb int = 2
	for !isPos {
		if isPrima(tempNumb) {
			pos++
		}
		if pos == number {
			break
		}
		// fmt.Println(tempNumb)
		tempNumb++
	}
	return tempNumb

}

func main() {
	// Problem 1 - Prima ke X (Tanpa Recrusive)
	fmt.Println(PrimetoX(1))
	fmt.Println(PrimetoX(5))
	fmt.Println(PrimetoX(8))
	fmt.Println(PrimetoX(9))
	fmt.Println(PrimetoX(10))

}
