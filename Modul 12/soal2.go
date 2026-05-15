package main

import "fmt"

func main() {
	var suara [21]int
	var input int
	var totalMasuk, suaraSah int

	for {
	fmt.Scan(&input)
	if input == 0 {
	break
		}
totalMasuk++

	if input >= 1 && input <= 20 {
	suara[input]++
	suaraSah++
	}
}

ketua, wakil := 0, 0
max1, max2 := -1, -1
	for i := 1; i <= 20; i++ {
		if suara[i] > max1 {
			max2 = max1
			wakil = ketua
			max1 = suara[i]
			ketua = i
		} else if suara[i] > max2 {
			max2 = suara[i]
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", suaraSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}