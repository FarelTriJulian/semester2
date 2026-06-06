package main

import "fmt"

func main() {
	var x string
	var n int

	fmt.Print("Masukkan string x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan jumlah data string: ")
	fmt.Scan(&n)

	dataString := make([]string, n)
	fmt.Printf("Masukkan %d data string:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("Data ke-%d: ", i+1)
		fmt.Scan(&dataString[i])
	}

	ketemu := false
	posisi := -1
	jumlahX := 0

	for i := 0; i < n; i++ {
		if dataString[i] == x {
			ketemu = true
			if posisi == -1 {
				posisi = i + 1
			}

			jumlahX++
		}
	}

	fmt.Println("\n--- HASIL ANALISIS ---")
	if ketemu {
		fmt.Println("a. Apakah string x ada? Ya, ada.")
	} else {
		fmt.Println("a. Apakah string x ada? Tidak ditemukan.")
	}

	if ketemu {
		fmt.Printf("b. Posisi pertama kali ditemukan: Ke-%d\n", posisi)
	} else {
		fmt.Println("b. Posisi pertama kali ditemukan: - (Tidak ada)")
	}

	fmt.Printf("c. Jumlah string x yang ditemukan: %d\n", jumlahX)

	if jumlahX >= 2 {
		fmt.Println("d. Adakah sedikitnya dua string x? Ya, ada.")
	} else {
		fmt.Println("d. Adakah sedikitnya dua string x? Tidak.")
	}
}
