package main

import (
	"fmt"
	"sort"
)

func GetStrongestKnight(dragon, knight []int) []int {
	sort.Ints(knight)
	// var j int

	// if len(dragon) > len(knight) {
	// 	j = len(dragon)
	// } else {
	// 	j = len(knight)
	// }

	// var theKnight []int = make([]int, j)
	// j = 0
	// // for i := len(knight) - 1; i >= 0; i-- {
	// for i := 0; i < len(knight); i++ {
	// 	theKnight[j] = knight[i]
	// 	j++
	// }
	// return theKnight
	return knight
}

func GetWeakestDragon(dragon, knight []int) []int {
	sort.Ints(dragon)
	// var j int

	// if len(dragon) > len(knight) {
	// 	j = len(dragon)
	// } else {
	// 	j = len(knight)
	// }

	// var theDragon []int = make([]int, j)
	// j = 0
	// for i := 0; i < len(dragon); i++ {
	// 	theDragon[j] = dragon[i]
	// 	j++
	// }
	// fmt.Println(theDragon)
	// return theDragon
	return dragon
}

func DragonOfLoowater(dragonHead, KnightHeight []int) {
	var countSuccess int
	// sumDragonHead := len(dragonHead)
	// fmt.Print("D", GetWeakestDragon(dragonHead, KnightHeight), "K", GetStrongestKnight(dragonHead, KnightHeight))
	StrongestKnight := GetStrongestKnight(dragonHead, KnightHeight)
	WeakestDragon := GetWeakestDragon(dragonHead, KnightHeight)

	// fmt.Println("D", WeakestDragon)
	// fmt.Println("K", StrongestKnight)

	var heightMinimum int
	// var sizeLoop int
	// if len(StrongestKnight) < len(WeakestDragon) {
	// 	sizeLoop = len(WeakestDragon)
	// } else {
	// 	sizeLoop = len(StrongestKnight)
	// }

	for i := 0; i < len(WeakestDragon); i++ {
		if WeakestDragon[i] == 0 {
			break
		}
		for j := 0; j < len(StrongestKnight); j++ {
			if StrongestKnight[j] == 0 {
				countSuccess--
				break
			}
			if WeakestDragon[i]-StrongestKnight[j] <= 0 {
				countSuccess++
				heightMinimum = heightMinimum + StrongestKnight[j]
				break

			}
		}
	}

	if countSuccess >= len(WeakestDragon) {
		// fmt.Println(countSuccess)
		fmt.Println(heightMinimum)
	} else {
		fmt.Println("knight fail")
	}
}

// func DragonOfLoowater(dragonHead, KnightHeight []int) {
// 	var minimumHeight, countSuccess, att int
// 	sumDragonHead := len(dragonHead)

// 	sort.Ints(KnightHeight)
// 	sort.Ints(dragonHead)

// 	if len(dragonHead) > len(KnightHeight) {
// 		att = len(KnightHeight)
// 	} else {
// 		att = len(dragonHead)

// 	}
// 	for i := len(KnightHeight); i >= att; i-- {
// 		// fmt.Println("D", dragonHead[sumDragonHead-1], "K", KnightHeight[i-1], "S", dragonHead[sumDragonHead-1]-KnightHeight[i-1])

// 		if sumDragonHead-1 < 0 {
// 			break
// 		}
// 		if dragonHead[sumDragonHead-1]-KnightHeight[i-1] < 0 {
// 			minimumHeight = minimumHeight + dragonHead[sumDragonHead-1] + 1
// 			countSuccess++
// 		}
// 		sumDragonHead--
// 	}
// 	if countSuccess >= len(dragonHead) {
// 		fmt.Println(minimumHeight)
// 	} else {
// 		fmt.Println("knight fail")
// 	}
// }

func main() {
	// Problem 3 - Dragon of Loowater
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})
	DragonOfLoowater([]int{5, 10}, []int{5})
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})
	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})
}
