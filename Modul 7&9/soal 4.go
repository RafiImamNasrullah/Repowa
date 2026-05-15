package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune


func isiArray(t *tabel, n *int) {
	var ch string
	*n = 0
	for *n < NMAX {
		fmt.Scan(&ch)
		if ch == "." {
			break
		}
		t[*n] = rune(ch[0])
		(*n)++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}


func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}


func palindrom(t tabel, n int) bool {

	var salin tabel = t
	balikanArray(&salin, n)

	for i := 0; i < n; i++ {
		if t[i] != salin[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks\t\t: ")
	isiArray(&tab, &m)

	isPalin := palindrom(tab, m)

	balikanArray(&tab, m)

	fmt.Print("Reverse teks\t: ")
	cetakArray(tab, m)

	fmt.Printf("Palindrom\t? %v\n", isPalin)
}