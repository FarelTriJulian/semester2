package main

import "fmt"

func ganjil(n int, i int) {
	if i > n {
		return
	}
	if i%2 == 1 {
		fmt.Print(i, " ")
	}
	ganjil(n, i+1)
}

func main() {
	var x int
	fmt.Print("Masukkan: ")
	fmt.Scan(&x)

	ganjil(x, 1)
}
