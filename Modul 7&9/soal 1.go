package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat titik
	r     int
}

func jarak(p, q titik) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var l1, l2 lingkaran
	var pt titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)

	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)

	fmt.Scan(&pt.x, &pt.y)

	d1 := didalam(l1, pt)
	d2 := didalam(l2, pt)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}