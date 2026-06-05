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

func main() {

	D := buatDominoes()
	kocokKartu(&D)

	fmt.Println("7 kartu pemain:")

	for i := 0; i < 7; i++ {
		k := ambilKartu(&D)
		fmt.Printf("[%d|%d] ", k.Kiri, k.Kanan)
	}
}