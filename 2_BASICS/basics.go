package main

import "fmt"

func main() {

	a := 5
	var b int = 4

	// if clause
	if a < 0 {
		fmt.Println("a is negative!")
	} else if a > b {
		fmt.Println("a is bigger than b")
	}

	//switch

	switch a {
	case 10:
		fmt.Println("It's 10")
	case 5:
		fmt.Print("It's five")
	}

	printArrayWithBasicForLooop()

	printArrayWithForRangeLoop()

	// printStringWithInfiniteLoop()

}

//lower case functions are private to the package!
func printArrayWithBasicForLooop() {
	// simple for loop
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
}

func printArrayWithForRangeLoop() {
	// for loop with range
	for _, i := range []int{1, 2, 3, 4} {
		fmt.Println(i)
	}
}

func printStringWithInfiniteLoop() {
	//infinite loop
	for {
		fmt.Println("this is infinite!")
	}
}
