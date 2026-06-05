package main

import "fmt"

func main() {

	var x float64
	var jumlah float64
	var rata float64
	var n int

	fmt.Println("Masukkan bilangan (akhiri dengan 9999):")

	fmt.Scan(&x)

	for x != 9999 {

		jumlah += x
		n++

		fmt.Scan(&x)
	}

	if n > 0 {
		rata = jumlah / float64(n)
		fmt.Println("Rata-rata =", rata)
	} else {
		fmt.Println("Tidak ada data")
	}
}