package main

import "fmt"

func main() {
	var tahun int

	fmt.Print("hahut: ")
	fmt.Scan(&tahun)

	Kabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	fmt.Printf("Kabisat: %t\n", Kabisat)
}
