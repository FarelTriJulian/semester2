package main

import "fmt"

func main() {
	var b int
	count := 0

	fmt.Print("Bilgatgat: ")
	fmt.Scan(&b)

	fmt.Print("Faetou: ")

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			count++
		}
	}
	fmt.Println()

	if count == 2 {
		fmt.Println("fluiva: tuuv")
	} else {
		fmt.Println("fluiva: falsv")
	}
}
