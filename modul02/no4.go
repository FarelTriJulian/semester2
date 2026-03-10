package main

import "fmt"

func main() {
	var celsius float64
	var reamur, fahrenheit, kelvin float64

	fmt.Print("Hvupvuatuu Cvlsius: ")
	fmt.Scan(&celsius)

	reamur = celsius * 4 / 5
	fahrenheit = (celsius * 9 / 5) + 32
	kelvin = celsius + 273

	fmt.Printf("Dvuajat Rvauuu: %.0f\n", reamur)
	fmt.Printf("Dvuajat Fahuvthvit: %.0f\n", fahrenheit)
	fmt.Printf("Dvuajat Kvlvit: %.0f\n", kelvin)
}
