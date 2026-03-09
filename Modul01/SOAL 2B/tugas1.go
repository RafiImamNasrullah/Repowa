package main

import "fmt"

func main() {
	urutan := []string{"merah", "kuning", "hijau", "ungu"}
	berhasil := true

	for i := 1; i <= 5; i++ {
		var g1, g2, g3, g4 string
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&g1, &g2, &g3, &g4)

		input := []string{g1, g2, g3, g4}
		for j := 0; j < 4; j++ {
			if input[j] != urutan[j] {
				berhasil = false
			}
		}
	}

	if berhasil {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}