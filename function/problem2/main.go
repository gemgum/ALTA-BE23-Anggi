package main

import "fmt"

func drawXYZ(angka int) {
	var spasi = angka
	var hasil int
	for i := 0; i < angka; i++ {
		for j := 0; j < spasi; j++ {
			hasil++
			switch {
			case hasil%3 == 0:
				{
					fmt.Print("X ")
				}
			case hasil%2 != 0:
				{
					fmt.Print("Y ")
				}
			case hasil%2 == 0:
				{
					fmt.Print("Z ")
				}

			}
		}

		fmt.Println("")

	}
}

func main() {
	// Problem 2 - Draw X Y Z
	var angka int
	fmt.Print("\nDraw X Y Z \n\n")
	fmt.Printf("Masukan Angka = ")
	fmt.Scanln(&angka)

	drawXYZ(angka)

}
