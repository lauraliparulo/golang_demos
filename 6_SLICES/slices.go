package main

import "fmt"

func main() {

	//variable lenght arrays
	slice := []int{1, 2, 3, 4}

	slice = append(slice, 4)

	fmt.Println(slice)

	//fixed lenght arrays
	fixed_slice := [4]int{1, 2, 3, 4}
	//cannot append another element - compiler error!
	//fixed_slice = append(slice, 4)

	fmt.Println(fixed_slice)

	//another way to define a slice
	another_slice := make([]int, 10)

	//initially filled with zeros
	fmt.Println(another_slice)

	//add a different number in the first position
	another_slice[0] = 45

	fmt.Println(another_slice)
}
