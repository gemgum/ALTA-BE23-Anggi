package main

import "fmt"

func RemoveDuplicate(arr []int) int {
	var temparr []int = arr
	var vArr []int
	for i := 0; i < len(arr); i++ {
		var unique bool = false
		for j := i + 1; j < len(temparr); j++ {
			if arr[i] == temparr[j] {
				unique = true
				break
			}
		}
		if !unique {
			vArr = append(vArr, arr[i])
		}
	}
	return len(vArr)
}

func main() {
	//Problem 5 - Remove Duplicates
	fmt.Println(RemoveDuplicate([]int{2, 3, 3, 6, 9, 9}))
	fmt.Println(RemoveDuplicate([]int{2, 3, 4, 5, 6, 9, 9}))
	fmt.Println(RemoveDuplicate([]int{2, 2, 2, 11}))
	fmt.Println(RemoveDuplicate([]int{2, 2, 2, 11}))
	fmt.Println(RemoveDuplicate([]int{1, 2, 3, 11, 11}))

}
