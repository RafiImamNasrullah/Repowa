package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func insertionSortAsc(T *arrInt, n int) {
	var temp, i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func main() {
	var data arrInt
	var n int
	n = 0

	var angka int
	fmt.Scan(&angka)

	for angka >= 0 {
		data[n] = angka
		n = n + 1
		fmt.Scan(&angka)
	}

	insertionSortAsc(&data, n)

	var k int
	k = 0
	for k < n {
		if k > 0 {
			fmt.Print(" ")
		}
		fmt.Print(data[k])
		k = k + 1
	}
	fmt.Println()

	if n <= 1 {
		fmt.Println("Data berjarak tidak tetap")
	} else {
		jarak := data[1] - data[0]
		tetap := true
		k = 1
		for k < n-1 {
			if data[k+1]-data[k] != jarak {
				tetap = false
			}
			k = k + 1
		}

		if tetap {
			fmt.Printf("Data berjarak %d\n", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}