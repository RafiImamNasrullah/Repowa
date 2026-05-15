package main

import "fmt"

const NMAX = 1000

type arrReal [NMAX]float64

func nilaiEkstrim(arr arrReal, n int) (min, max float64) {
	min = arr[0]
	max = arr[0]

	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	return min, max
}

func main() {
	var arr arrReal
	var n int

	fmt.Print("Masukkan jumlah anak kelinci (N): ")
	fmt.Scan(&n)

	fmt.Println("Masukkan berat tiap anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Printf("  Kelinci ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	min, max := nilaiEkstrim(arr, n)

	fmt.Printf("\nBerat terkecil : %.2f\n", min)
	fmt.Printf("Berat terbesar : %.2f\n", max)
}