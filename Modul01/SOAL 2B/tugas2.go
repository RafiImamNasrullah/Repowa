package main

import "fmt"

func main() {
	var n int
	var bunga string
	var pita string
	var jumlah int

	fmt.Print("N: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print("Bunga ", i, ": ")
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita = pita + bunga + " - "
		jumlah++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}