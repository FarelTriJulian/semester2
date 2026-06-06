package main

import "fmt"

func main() {
	var bilangan, total float64
	var jumlahdata int

	fmt.Println("Masukkan sejumlah bilangan real (akhiri dengan 9999):")

	for {
		fmt.Scan(&bilangan)
		if bilangan == 9999 {
			break
		}
		total += bilangan
		jumlahdata++
	}

	if jumlahdata > 0 {
		rerata := total / float64(jumlahdata)
		fmt.Printf("\nRata-rata dari bilangan tersebut adalah: %.2f\n", rerata)
	} else {
		fmt.Println("\nTidak ada data bilangan yang valid untuk dihitung.")
	}
}
