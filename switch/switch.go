package main

import (
	"fmt"
	"time"
)

func main() {
	m := 9
	switch m {
	case 5:
		fmt.Println("five")
		fmt.Println("six")
	case 10:
		fmt.Println("10..")
		fmt.Println("11..")
	default:
		fmt.Println("nothing")
	}

	switch {
	case m < 5, m < 10:
		fmt.Println("m < 10")
	case m == 10:
		fmt.Println("m == 10")
	default:
		fmt.Println("done")

	}

	mytype := func(t interface{}) {
		switch ty := t.(type) {
		case bool:
			fmt.Println("bool", ty)
		case int:
			fmt.Println("int", ty)
		case float64:
			fmt.Println("float64", ty)
		case string:
			fmt.Println("string", ty)
		}
	}

	mytype("hello")
	mytype(1)
	mytype(true)
	mytype(1.2)

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() >= 12:
		fmt.Println("Evening")
	}
}
