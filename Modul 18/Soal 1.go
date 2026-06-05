package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	Kiri  int
	Kanan int
	Balak bool
}

type Dominoes struct {
	Kartu [28]Domino
	Sisa  int
}

func buatDominoes() Dominoes {
	var D Dominoes
	idx := 0

	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			D.Kartu[idx] = Domino{i, j, i == j}
			idx++
		}
	}

	D.Sisa = 28
	return D
}

func kocokKartu(D *Dominoes) {
	rand.Seed(time.Now().UnixNano())

	for i := D.Sisa - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		D.Kartu[i], D.Kartu[j] = D.Kartu[j], D.Kartu[i]
	}
}

func ambilKartu(D *Dominoes) Domino {
	D.Sisa--
	return D.Kartu[D.Sisa]
}

func gambarKartu(d Domino, suit int) int {
	if suit == 1 {
		return d.Kiri
	}
	return d.Kanan
}

func nilaiKartu(d Domino) int {
	return d.Kiri + d.Kanan
}

func main() {
	D := buatDominoes()
	kocokKartu(&D)

	k := ambilKartu(&D)

	fmt.Println("Kartu :", k)
	fmt.Println("Gambar kiri :", gambarKartu(k, 1))
	fmt.Println("Gambar kanan :", gambarKartu(k, 2))
	fmt.Println("Nilai :", nilaiKartu(k))
}