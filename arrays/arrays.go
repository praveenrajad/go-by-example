package main

import "fmt"

func main() {
	var a [5] int
	fmt.Println (a)
	a[3] = 100
	fmt.Println(a)

	x := [5][5]int{}
	fmt.Println(x)
	for i:=0; i < 5; i++ {
		for j:=0; j < 5; j++ {
			x[i][j] = i + j
		}
	} 
	fmt.Println(x)
	fmt.Println(len(x), cap(x))
}