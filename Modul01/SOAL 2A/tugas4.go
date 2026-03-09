package main

import "fmt"

func main() {
	var celsius float64
	var reamur, fahrenheit, kelvin float64

	fmt.Print("Input Celsius: ")
	fmt.Scanln(&celsius)

	reamur = celsius * 4 / 5
	fahrenheit = celsius*9/5 + 32
	kelvin = celsius + 273

	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}