package main

import "fmt"

func main() {
	var tahunlahir int
	var tahunsekarang int

	fmt.Println("Masukkan tahun lahirmu dan tahun ini! ex:(2006 2025)")
	fmt.Scan(&tahunlahir, &tahunsekarang)

	usia := tahunsekarang - tahunlahir

	fmt.Printf("Umurmu %d tahun. \n", usia)
}
