package main

import "fmt"

func main() {
	var klubA, klubB string

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	var skorA, skorB int
	p := 1

	for {
		fmt.Printf("Pertandingan %d: ", p)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}
		if skorA > skorB {
			fmt.Printf("Hasil %d : %s\n", p, klubA)
		} else if skorB > skorA {
			fmt.Printf("Hasil %d : %s\n", p, klubB)
		} else {
			fmt.Printf("Hasil %d : Draw\n", p)
		}
		p++
	}
}
