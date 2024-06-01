package main

import "fmt"

func moneyChange(money int) []int {
	var pecahan []int = []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	var data []int
	for i := len(pecahan) - 1; i >= 0; i-- {
		if money-pecahan[i] >= 0 {
			for money >= pecahan[i] {
				data = append(data, pecahan[i])
				money = money - pecahan[i]
			}
		}
	}

	return data
}

func main() {
	// Problem 2 - Min and Max Using Pointer
	fmt.Println(moneyChange(123))
	fmt.Println(moneyChange(432))
	fmt.Println(moneyChange(543))
	fmt.Println(moneyChange(7752))
	fmt.Println(moneyChange(15321))
}
