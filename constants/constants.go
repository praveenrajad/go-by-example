package main

import (
	"fmt"
	"math"
)

func main() {
	const i = 1000
	const (
		x = "hello"
		a = 10
	)
	test := 3e2 / a
	fmt.Println(int(test))
	test2 := math.Sin(a)
	fmt.Println(test2)
	const z int = 100
	test3 := 3e2/z
	fmt.Println(test3)
	// n := 1.2
	// fmt.Println(z/n)
}
