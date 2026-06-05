package main

import "fmt"

func main() {

	var x string
	var n int

	fmt.Print("Masukkan string x: ")
	fmt.Scan(&x)

	fmt.Print("Jumlah data string: ")
	fmt.Scan(&n)

	data := make([]string, n)

	fmt.Println("Masukkan data string:")

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	ketemu := false


	posisi := -1
	jumlah := 0

	for i := 0; i < n; i++ {

		if data[i] == x {

			jumlah++

			if !ketemu {
				ketemu = true
				posisi = i + 1
			}
		}
	}
	if ketemu {
		fmt.Println("String ditemukan")
	} else {
		fmt.Println("String tidak ditemukan")
	}
	if posisi != -1 {
		fmt.Println("Posisi pertama:", posisi)
	}
	fmt.Println("Jumlah string x:", jumlah)
	if jumlah >= 2 {
		fmt.Println("Terdapat minimal dua string x")
	} else {
		fmt.Println("Tidak sampai dua")
	}
}