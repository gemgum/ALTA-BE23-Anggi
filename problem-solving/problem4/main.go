package main

import "fmt"

func searchRight(arr []int, val int) int {
	var rv int = -1
	for i := len(arr) - 1; i >= len(arr)/2; i-- {
		if arr[i] == val {
			rv = i
		}
	}
	return rv
}

func searchLeft(arr []int, val int) int {
	var rv int = -1
	for i := len(arr) / 2; i >= 0; i-- {
		if arr[i] == val {
			// fmt.Println("L", arr[i])
			rv = i
		}
	}
	return rv
}
func BinarySearch(array []int, x int) {
	var result int
	midVal := len(array) / 2
	if array[midVal] > x {
		result = searchLeft(array, x)
	} else {
		result = searchRight(array, x)
	}
	fmt.Println(result)
}

func main() {
	// Problem 4 - Binary Search Algoritm

	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100)

}
