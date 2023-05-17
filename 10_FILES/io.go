package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	err := ioutil.WriteFile("output.txt", []byte("Hello, world!"), 0644)

	check(err)

	// file2, err := os.Create("file2.txt")
	// defer file2.Close()

	buffer, err2 := ioutil.ReadFile("output.txt")

	check(err2)
	fmt.Println(buffer)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
