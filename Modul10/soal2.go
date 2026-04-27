package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var berat [1000]float64

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := x / y
	var totalWadah [1000]float64
	var totalSemua float64

	index := 0

	for i := 0; i < jumlahWadah; i++ {
		total := 0.0
		for j := 0; j < y; j++ {
			total += berat[index]
			index++
		}
		totalWadah[i] = total
		totalSemua += total
	}
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f ", totalWadah[i])
	}
	fmt.Println()
	rataRata := totalSemua / float64(jumlahWadah)
	fmt.Printf("%.2f\n", rataRata)
}
