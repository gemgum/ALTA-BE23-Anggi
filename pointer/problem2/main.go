package main

import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {
	max = *numbers[0]
	min = *numbers[0]

	for _, v := range numbers {
		if *v > max {
			max = *v
		}
		if *v < min {
			min = *v
		}
	}
	return min, max
}

func main() {
	// Problem 2 - Min and Max Using Pointer
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}
