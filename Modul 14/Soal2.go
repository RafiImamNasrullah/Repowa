package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		var jumlah int
		fmt.Scan(&jumlah)

		var angka int
		var ganjil [100]int
		var genap [100]int
		var g, e int

		for j := 0; j < jumlah; j++ {
			fmt.Scan(&angka)

			if angka%2 == 1 {
				ganjil[g] = angka
				g++
			} else {
				genap[e] = angka
				e++
			}
		}

		for a := 0; a < g-1; a++ {
			for b := a + 1; b < g; b++ {
				if ganjil[a] > ganjil[b] {
					ganjil[a], ganjil[b] = ganjil[b], ganjil[a]
				}
			}
		}

		for a := 0; a < e-1; a++ {
			for b := a + 1; b < e; b++ {
				if genap[a] < genap[b] {
					genap[a], genap[b] = genap[b], genap[a]
				}
			}
		}

		for j := 0; j < g; j++ {
			fmt.Print(ganjil[j], " ")
		}

		for j := 0; j < e; j++ {
			fmt.Print(genap[j], " ")
		}

		fmt.Println()
	}
}