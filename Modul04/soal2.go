package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var ts, tw int
	var max, min int

	max = -1
	min = 999999

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		hitungSkor(&ts, &tw)
		if ts > max {
			max = ts
			min = tw
			pemenang = nama
		} else if ts == max {
			if tw < min {
				min = tw
				pemenang = nama
			}
		}
	}

	if pemenang != "" {
		fmt.Printf("%s %d %d\n", pemenang, max, min)
	}
}
