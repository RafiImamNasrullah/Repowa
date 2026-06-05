package main

import (
	"fmt"
)

var data string
var pos int

func start(teks string) {
	data = teks
	pos = 0
}

func maju() {
	if pos < len(data)-1 {
		pos++
	}
}

func eop() bool {
	return data[pos] == '.'
}

func cc() byte {
	return data[pos]
}

func hitungKarakter() int {
	jumlah := 0

	for !eop() {
		jumlah++
		maju()
	}

	return jumlah
}

func hitungA() int {
	jumlah := 0

	for !eop() {
		if cc() == 'A' {
			jumlah++
		}
		maju()
	}

	return jumlah
}

func main() {

	start("SAYA BELAJAR ALGORITMA.")

	fmt.Println("Jumlah karakter :", hitungKarakter())

	start("SAYA BELAJAR ALGORITMA.")
	fmt.Println("Jumlah huruf A :", hitungA())
}