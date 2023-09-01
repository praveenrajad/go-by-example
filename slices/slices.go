package main

import (
	"fmt"
)

func main() {
	{
		x := make([]string, 5)
		y := make([]string, 5, 10)
		var z []string
		x1 := []string{}

		fmt.Println(x, y, z, x1)
		fmt.Println(len(x), len(y), len(z), len(x1))
		fmt.Println(cap(x), cap(y), cap(z), cap(x1))

	}
	{
		x := []int{1, 2, 3, 4, 5}
		y := make([]int, len(x))
		fmt.Println(x)
		copy(y, x)
		fmt.Println(y)
		x[0] = 10
		fmt.Println(x, y)

		z := x
		x[1] = 20
		fmt.Println(z)
		x = append(x, 100, 200)
		x[2] = 30
		fmt.Println(x)
		fmt.Println(z)
	}
	{
		x := []int{1, 2, 3, 4, 5}
		y := x[3:]
		fmt.Println(x, y)
		x[3] = 40
		fmt.Println(x, y)
	}

	{
		arr := make([][]int, 5) 
		for i:= 0; i < 5; i++ {
			arr[i] = make([]int, i+1)
			for j:=0; j < i+1; j++ {
				arr[i][j] = 1
			}
		}
		fmt.Println(arr)
	}
}
