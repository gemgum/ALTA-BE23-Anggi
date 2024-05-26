package main

import (
	"fmt"
)

func findMaxSumSubArray(maxIndex int, arr []int) int {
	var rValue int
	for ind, _ := range arr {
		var sumVal int
		for i := 0; i < maxIndex; i++ {
			if (i + ind) >= len(arr) {
				break
			}
			sumVal = sumVal + arr[i+ind]
		}
		if sumVal > rValue {
			rValue = sumVal
		}
	}
	return rValue
}
func main() {
	// Problem 4 - Maximum Sum Subarray of Size K

	fmt.Println(findMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2}))
	fmt.Println(findMaxSumSubArray(2, []int{2, 3, 4, 1, 5}))
	fmt.Println(findMaxSumSubArray(2, []int{2, 1, 4, 1, 1}))
	fmt.Println(findMaxSumSubArray(3, []int{2, 1, 4, 1, 1}))
	fmt.Println(findMaxSumSubArray(4, []int{2, 1, 4, 1, 1}))

}
