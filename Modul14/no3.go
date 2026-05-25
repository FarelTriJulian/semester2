package main

import "fmt"

const NMAX int = 100000

type arrInt [NMAX]int

func insertionSort(A *arrInt, n int) {
	for i := 1; i < n; i++ {
		temp := A[i]
		j := i - 1

		for j >= 0 && A[j] > temp {
			A[j+1] = A[j]
			j--
		}

		A[j+1] = temp
	}
}

func median(A arrInt, n int) int {
	if n%2 == 1 {
		return A[n/2]
	}
	return (A[(n/2)-1] + A[n/2]) / 2
}

func main() {
	var A arrInt
	var x, n int

	n = 0

	for {
		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {
			insertionSort(&A, n)
			fmt.Println(median(A, n))
		} else {
			A[n] = x
			n++
		}
	}
}
