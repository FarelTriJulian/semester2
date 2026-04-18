package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}
type Circle struct {
	center Point
	radius float64
}

func distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}
func isInside(c Circle, p Point) bool {
	return distance(c.center, p) <= c.radius
}
func main() {
	var c1, c2 Circle
	var p Point

	fmt.Scan(&c1.center.x, &c1.center.y, &c1.radius)
	fmt.Scan(&c2.center.x, &c2.center.y, &c2.radius)
	fmt.Scan(&p.x, &p.y)

	inC1 := isInside(c1, p)
	inC2 := isInside(c2, p)

	if inC1 && inC2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inC1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inC2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
