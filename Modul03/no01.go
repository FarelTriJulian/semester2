package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result
}
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}
func main() {
	var a, b, c, d int

	fmt.Print("masukkan empat bilangan: ")
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutation(a, c)
	c1 := combination(a, c)

	p2 := permutation(b, d)
	c2 := combination(b, d)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
