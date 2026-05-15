package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	var hasil [100]string
	jumlah := 0
	pertandingan := 1

	for {

		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[jumlah] = klubA
		} else if skorB > skorA {
			hasil[jumlah] = klubB
		} else {
			hasil[jumlah] = "Draw"
		}

		jumlah++
		pertandingan++
	}

	for i := 0; i < jumlah; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}

	fmt.Println("Pertandingan selesai")
}