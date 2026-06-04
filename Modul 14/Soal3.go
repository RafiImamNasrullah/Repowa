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

	for angka != -5313 {
		if angka == 0 {
			insertionSortAsc(&data, n)

			var median int
			if n%2 != 0 {
				median = data[n/2]
			} else {
				median = (data[n/2-1] + data[n/2]) / 2
			}
			fmt.Println(median)
		} else {
			data[n] = angka
			n = n + 1
		}
		fmt.Scan(&angka)
	}
}