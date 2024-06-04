package main

import "fmt"

func recFibonacci(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		// fmt.Println("F1", recFibonacci(number-1))
		// fmt.Println("F2", recFibonacci(number-2))
		// fmt.Println("RV", recFibonacci(number-1)+recFibonacci(number-2))

		return recFibonacci(number-1) + recFibonacci(number-2) // F1-1 + F2-2 -> (3 - 1) + (2 - 2) -> 2 + 0  = 1
		//				return recFibonacci(number-1) + recFibonacci(number-2) // F1-1 + F2-2 -> (2 - 1) + (2 - 2) -> 1 + 0  = 1

	}
}

// F4 + F5 -> (3) + (5) -> 3 + 5  = 8  ->F6
// F3 + F4 -> (2) + (3) -> 2 + 3  = 5  ->F5
// F2 + F3 -> (1) + (2) -> 1 + 2  = 3  ->F4
// F2 + F1 -> (1) + (1) -> 1 + 1  = 2  ->F3
// F1 + F0 -> (1) + (0) -> 1 + 0  = 1  ->F2

// func Fibonacci(number int) int {
// 	var t1, t2 int = 0, 1
// 	var rv int
// 	for i := 1; i < number; i++ {
// 		rv = t1 + t2
// 		t1 = t2
// 		t2 = rv
// 	}
// 	return rv
// }

func main() {
	// Problem 2 - Recrusive Fibonacci
	fmt.Println(recFibonacci(0))
	fmt.Println(recFibonacci(2))
	fmt.Println(recFibonacci(9))
	fmt.Println(recFibonacci(10))
	fmt.Println(recFibonacci(12))

	// fmt.Println(Fibonacci(0))
	// fmt.Println(Fibonacci(2))
	// fmt.Println(Fibonacci(9))
	// fmt.Println(Fibonacci(10))
	// fmt.Println(Fibonacci(12))
}
