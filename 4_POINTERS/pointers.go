package main

import (
	"fmt"
)

func main() {
	s := "this is a string"

	string_pointer := &s

	fmt.Println(s)

	//reference the pointer with asteryx
	fmt.Println("reference by pointer: ", *string_pointer)
	//reference the pointer with asteryx
	fmt.Println("address: ", string_pointer)
}
