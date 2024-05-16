package main

import "fmt"

func main() {
	//Problem 6 - Program untuk menghitung harga setelah diskon
	var hargaBarangAwal float32
	var diskon float32
	var potonganDiskon float32

	fmt.Print("\nPenghitung Diskon Barang\n\n")
	fmt.Printf("Masukan Harga Barang = Rp.")
	fmt.Scanln(&hargaBarangAwal)

	fmt.Printf("Masukan Diskon = ")
	fmt.Scanln(&diskon)

	potonganDiskon = (diskon / 100) * hargaBarangAwal

	hasil := hargaBarangAwal - potonganDiskon
	fmt.Printf("\nMaka Harga yang harus dibayar = Rp.%.2f\n", hasil)

}
