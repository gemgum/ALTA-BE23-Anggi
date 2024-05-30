package main

import "fmt"

func swap(a, b *int) {
	var c int = *b
	*b = *a
	*a = c

}

func main() {
	// Problem 1 - Swap Two Number Using Pointer
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println(a, b)
}
