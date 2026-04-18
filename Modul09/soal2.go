package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nSeluruh isi array:", arr)

	fmt.Print("Elemen indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Elemen indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Printf("Elemen indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idx int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)

	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println("Array setelah penghapusan:", arr)

	var sum int
	for _, v := range arr {
		sum += v
	}
	avg := float64(sum) / float64(len(arr))
	fmt.Println("Rata-rata:", avg)

	var variance float64
	for _, v := range arr {
		variance += math.Pow(float64(v)-avg, 2)
	}
	variance /= float64(len(arr))
	stdDev := math.Sqrt(variance)
	fmt.Println("Standar deviasi:", stdDev)

	var target, freq int
	fmt.Print("Masukkan angka yang ingin dicari frekuensinya: ")
	fmt.Scan(&target)

	for _, v := range arr {
		if v == target {
			freq++
		}
	}
	fmt.Printf("Frekuensi %d: %d\n", target, freq)
}
