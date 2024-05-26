package main

import "fmt"

func ArrayUniqueOld(arrA, arrB []int) []int {
	vArr := make([]int, 0)
	for i := 0; i < len(arrA); i++ {
		var unique bool = false
		for j := 0; j < len(arrB); j++ {
			if arrA[i] == arrB[j] {
				unique = true
				break
			}
		}
		if !unique {
			vArr = append(vArr, arrA[i])
		}
	}
	return vArr
}

// func ArrayUniqueNew(arrA, arrB []int) []int {
// 	vArr := make([]int, 0)
// 	for _, a := range arrA {
// 		var unique bool = false
// 		for _, b := range arrB {
// 			if a == b {
// 				unique = true
// 				break
// 			}
// 		}
// 		if !unique {
// 			vArr = append(vArr, a)
// 		}
// 	}
// 	return vArr
// }
func main() {
	// Problem 3 - Array Unique
	fmt.Println(ArrayUniqueOld([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16}))
	fmt.Println(ArrayUniqueOld([]int{10, 20, 30, 40}, []int{5, 10, 15, 59}))
	fmt.Println(ArrayUniqueOld([]int{1, 3, 7}, []int{1, 3, 5}))
	fmt.Println(ArrayUniqueOld([]int{3, 8}, []int{2, 8}))
	fmt.Println(ArrayUniqueOld([]int{1, 2, 3}, []int{3, 2, 1}))

}
