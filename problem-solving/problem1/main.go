package main

import "fmt"

func SimpleEqualations(x, y, z int) {
	// higher := x * y * z

	// for i := 0; i < higher; i++ {
	// 	for j := 0; j < higher; j++ {
	// 		for k := 0; k < higher; k++ {
	// 			if i+j+k == x && i*j*k == y && (i*i)+(j*j)+(k*k) == z {
	// 				fmt.Println(i, j, k)
	// 				return
	// 			}
	// 		}
	// 	}
	// }

	for i := 1; i < x; i++ {
		for j := i; j < x-i; j++ {
			k := x - i - j
			if k > 0 && i*j*k == y && i*i+j*j+k*k == z {
				fmt.Println(i, j, k)
				return
			}
		}
	}

	fmt.Println("no solution")

}

func main() {
	// Problem 1 - Simple Equalations
	SimpleEqualations(9, 24, 29)
}
