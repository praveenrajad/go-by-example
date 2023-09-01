package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i == 1 {
			continue
		}

		fmt.Println("inf", i) 
		if i == 2 {
			break
		}
	}


	for i:=0; i < 5; i ++ {
		fmt.Println("Normal", i)
	} 

}