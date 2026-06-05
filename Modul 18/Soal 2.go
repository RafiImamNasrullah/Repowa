package main

import (
	"fmt"
)

type Domino struct {
	Kiri  int
	Kanan int
	Balak bool
}

func nilaiKartu(d Domino) int {
	return d.Kiri + d.Kanan
}

func sepasangKartu(a, b Domino) bool {
	return nilaiKartu(a)+nilaiKartu(b) == 12
}

func main() {

	k1 := Domino{6, 3, false}
	k2 := Domino{2, 1, false}

	fmt.Println("Kartu 1 :", k1)
	fmt.Println("Kartu 2 :", k2)

	fmt.Println("Sepasang :", sepasangKartu(k1, k2))
}