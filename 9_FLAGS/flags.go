package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var name string
	flag.StringVar(&name, "name", "", "The name to create the hello world message")
	flag.Parse()

	if name == "" {
		fmt.Println("Enter your name:")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Hello,", name)
}
