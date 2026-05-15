package main

import (
	"fmt"
	"math"
)

const NMAX = 100

type arrInt [NMAX]int

func tampilSemua(arr arrInt, n int) {
	fmt.Print("Isi array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("[%d]=%d ", i, arr[i])
	}
	fmt.Println()
}

func tampilGanjil(arr arrInt, n int) {
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Printf("[%d]=%d ", i, arr[i])
	}
	fmt.Println()
}

func tampilGenap(arr arrInt, n int) {
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Printf("[%d]=%d ", i, arr[i])
	}
	fmt.Println()
}

func tampilKelipatanX(arr arrInt, n, x int) {
	fmt.Printf("Indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if x != 0 && i%x == 0 {
			fmt.Printf("[%d]=%d ", i, arr[i])
		}
	}
	fmt.Println()
}

func hapusElemen(arr *arrInt, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[*n-1] = 0 
	*n--
}

func rataRata(arr arrInt, n int) float64 {
	if n == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(n)
}

func standarDeviasi(arr arrInt, n int) float64 {
	if n == 0 {
		return 0
	}
	rata := rataRata(arr, n)
	sumKuadrat := 0.0
	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rata
		sumKuadrat += selisih * selisih
	}
	return math.Sqrt(sumKuadrat / float64(n))
}

func frekuensi(arr arrInt, n, bil int) int {
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] == bil {
			count++
		}
	}
	return count
}

func main() {
	var arr arrInt
	var n int

	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("  arr[%d] = ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\n========== HASIL ==========")

	tampilSemua(arr, n)

	tampilGanjil(arr, n)

	tampilGenap(arr, n)

	var x int
	fmt.Print("\nMasukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	tampilKelipatanX(arr, n, x)

	var idx int
	fmt.Print("\nMasukkan indeks elemen yang ingin dihapus: ")
	fmt.Scan(&idx)
	if idx >= 0 && idx < n {
		hapusElemen(&arr, &n, idx)
		fmt.Printf("Array setelah menghapus indeks %d:\n", idx)
		tampilSemua(arr, n)
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	fmt.Printf("\nRata-rata        : %.4f\n", rataRata(arr, n))

	fmt.Printf("Standar deviasi  : %.4f\n", standarDeviasi(arr, n))

	var bil int
	fmt.Print("\nMasukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&bil)
	fmt.Printf("Frekuensi %d dalam array: %d kali\n", bil, frekuensi(arr, n, bil))
}