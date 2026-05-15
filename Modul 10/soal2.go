package main

import "fmt"

const NMAX = 1000

func main() {
	var x, y int
	var ikan [NMAX]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	jumlahWadah := (x + y - 1) / y
	var totalSemua float64

	for i := 0; i < jumlahWadah; i++ {
		var totalWadah float64

		for j := i * y; j < (i+1)*y && j < x; j++ {
			totalWadah += ikan[j]
		}

		fmt.Printf("%.2f ", totalWadah)
		totalSemua += totalWadah
	}

	rata := totalSemua / float64(jumlahWadah)

	fmt.Printf("\n%.2f\n", rata)
}