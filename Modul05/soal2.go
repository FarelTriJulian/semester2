package main

import "fmt"

func cb(n int, current int) {
	if current > n {
		return
	}

	for i := 0; i < current; i++ {
		fmt.Print("*")
	}
	fmt.Println()

	cb(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&n)

	cb(n, 1)
}
