package main

import "fmt"

func MaxSequance(arr []int) int {

	maxSoFar := arr[0]
	maxEndingHere := arr[0]

	for i := 1; i < len(arr); i++ {
		if maxEndingHere < 0 {
			maxEndingHere = arr[i]
		} else {
			maxEndingHere += arr[i]
		}

		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar

}
func main() {
	// Problem 4 - Binary Search Algoritm
	fmt.Println(MaxSequance([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequance([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequance([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequance([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequance([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
