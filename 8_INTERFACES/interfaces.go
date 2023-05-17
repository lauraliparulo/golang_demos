package main

import "fmt"

func main() {

	//empty interface
	var x interface{}
	x = "Hello, World"
	s, ok := x.(string)

	if !ok {
		panic("no int!")
	} else {
		fmt.Println("s", s)
	}

	switch x.(type) {
	case int:
		fmt.Println("Integer!")
	case string:
		fmt.Println("String!")
	}
}
