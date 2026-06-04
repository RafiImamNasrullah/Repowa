package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func selectionSort(T *arrInt, n int) {
	var t, i, j, idxMin int
	i = 1
	for i <= n-1 {
		idxMin = i - 1
		j = i
		for j < n {
			if T[idxMin] > T[j] {
				idxMin = j
			}
			j = j + 1
		}
		t = T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n, m int
	fmt.Scan(&n)

	var daerah int
	daerah = 0
	for daerah < n {
		fmt.Scan(&m)
		var rumah arrInt
		var k int
		k = 0
		for k < m {
			fmt.Scan(&rumah[k])
			k = k + 1
		}

		selectionSort(&rumah, m)

		k = 0
		for k < m {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[k])
			k = k + 1
		}
		fmt.Println()
		daerah = daerah + 1
	}
}