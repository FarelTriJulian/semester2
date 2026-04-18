package main

import "fmt"

const MAX = 127

type tabChar struct {
	tab [MAX]rune
	n   int
}

func isiArray(t *tabChar) {
	var input string
	fmt.Print("Teks: ")
	fmt.Scan(&input)

	t.n = len(input)
	for i := 0; i < t.n; i++ {
		t.tab[i] = rune(input[i])
	}
}
func cetakArray(t tabChar) {
	for i := 0; i < t.n; i++ {
		fmt.Printf("%c ", t.tab[i])
	}
	fmt.Println()
}
func balikArray(t tabChar) tabChar {
	var hasil tabChar
	hasil.n = t.n

	for i := 0; i < t.n; i++ {
		hasil.tab[i] = t.tab[t.n-1-i]
	}
	return hasil
}
func palindrome(t tabChar) bool {
	for i := 0; i < t.n/2; i++ {
		if t.tab[i] != t.tab[t.n-1-i] {
			return false
		}
	}
	return true
}
func main() {
	var t tabChar

	isiArray(&t)
	fmt.Print("Teks: ")
	cetakArray(t)

	rev := balikArray(t)
	fmt.Print("Reverse teks: ")
	cetakArray(rev)

	if palindrome(t) {
		fmt.Println("Palindrome: true")
	} else {
		fmt.Println("Palindrome: false")
	}
}
