package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	delete(m, "one")
	fmt.Println(m)
	m = make(map[string]int)
	fmt.Println(m)
	m["three"] = 3
	val, ok := m["two"]
	if !ok {
		fmt.Println("Not found")
	}
	fmt.Println(val)

	m = map[string]int{"four": 4}
	fmt.Println(m)
}
