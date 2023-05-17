package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("output.txt")

	if err != nil {
		panic("unable to create the file")
	}

	defer file.Close()

	file.WriteString("Hello laura")

	//read in the byte buffer
	buffer := make([]byte, 64)

	countBytes, err := file.Read(buffer)

	if err != nil {
		panic("unable to read the file")
	}

	fmt.Printf("Read all ", countBytes, " bytes")

	fmt.Println(buffer)

	fmt.Printf("Read non empty", countBytes, " bytes")

	fmt.Println(buffer[:countBytes])
	//with string cast
	fmt.Println(string(buffer[:countBytes]))

}
