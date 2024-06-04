package main

import "fmt"

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

func SquaePrime(wide, high, start int) string {
	var rv string

	// rv = rv + fmt.Sprint("\n")
	k := start
	var sumPrime int
	for i := 0; i < high; i++ {
		for j := 0; j < wide; j++ {
			var tempIsPrima bool = false
			for !tempIsPrima {
				k++
				tempIsPrima = isPrima(k)
				// fmt.Println(k)
				if tempIsPrima {
					rv = rv + fmt.Sprint(k) + " "
					sumPrime = sumPrime + k
				}
			}

		}
		rv = rv + "\n"
	}
	// rv = rv + "\n"
	rv = rv + fmt.Sprint(sumPrime) + " "
	return rv
}

func main() {
	// Problem 3 - Prima Segi Empat (Tanpa Recrusive)
	fmt.Println(SquaePrime(2, 3, 13))
	fmt.Println(SquaePrime(5, 2, 1))
}
