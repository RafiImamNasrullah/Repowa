package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	fmt.Print("Banyak Topping: ")
	fmt.Scan(&n)

	onPizza := 0
	cx, cy := 0.5, 0.5
	r := 0.5

	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		dx := x - cx
		dy := y - cy
		if dx*dx+dy*dy <= r*r {
			onPizza++
		}
	}

	fmt.Printf("Topping pada Pizza: %d\n", onPizza)

	pi := 4.0 * float64(onPizza) / float64(n)
	fmt.Printf("PI : %.10f\n", pi)
}