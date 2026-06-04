package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id       string
	judul    string
	penulis  string
	penerbit string
	eksemplar int
	tahun    int
	rating   int
}

type DaftarBuku [nMax]Buku

var pustaka DaftarBuku
var nPustaka int

func DaftarkanBuku(n int) {
	var i int
	i = 0
	for i < n {
		fmt.Scan(&pustaka[i].id)
		fmt.Scan(&pustaka[i].judul)
		fmt.Scan(&pustaka[i].penulis)
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Scan(&pustaka[i].tahun)
		fmt.Scan(&pustaka[i].rating)
		i = i + 1
	}
	nPustaka = n
}

func CetakTerfavorit(n int) {
	var idxMax, i int
	idxMax = 0
	i = 1
	for i < n {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
		i = i + 1
	}
	fmt.Println("Buku Terfavorit:")
	fmt.Println("Judul    :", pustaka[idxMax].judul)
	fmt.Println("Penulis  :", pustaka[idxMax].penulis)
	fmt.Println("Penerbit :", pustaka[idxMax].penerbit)
	fmt.Println("Tahun    :", pustaka[idxMax].tahun)
}

func UrutBuku(n int) {
	var temp Buku
	var i, j int
	i = 1
	for i <= n-1 {
		j = i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j = j - 1
		}
		pustaka[j] = temp
		i = i + 1
	}
}

func Cetak5Terbaru(n int) {
	fmt.Println("\n5 Buku dengan Rating Tertinggi:")
	limit := 5
	if n < 5 {
		limit = n
	}
	var i int
	i = 0
	for i < limit {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
		i = i + 1
	}
}

func CariBuku(n int, r int) {
	var lo, hi, mid int
	lo = 0
	hi = n - 1
	var found bool
	found = false
	var idxFound int
	idxFound = -1

	for lo <= hi {
		mid = (lo + hi) / 2
		if pustaka[mid].rating == r {
			found = true
			idxFound = mid
			lo = hi + 1 
		} else if pustaka[mid].rating < r {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	fmt.Println("\nHasil Pencarian Buku:")
	if found {
		fmt.Println("Judul     :", pustaka[idxFound].judul)
		fmt.Println("Penulis   :", pustaka[idxFound].penulis)
		fmt.Println("Penerbit  :", pustaka[idxFound].penerbit)
		fmt.Println("Tahun     :", pustaka[idxFound].tahun)
		fmt.Println("Eksemplar :", pustaka[idxFound].eksemplar)
		fmt.Println("Rating    :", pustaka[idxFound].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var n, r int
	fmt.Scan(&n)

	DaftarkanBuku(n)
	CetakTerfavorit(n)
	UrutBuku(n)
	Cetak5Terbaru(n)
	fmt.Scan(&r)
	CariBuku(n, r)
}