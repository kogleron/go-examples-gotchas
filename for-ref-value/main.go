package main

import "fmt"

//Expects 1,2,3,4,5
func main() {
	source := []int{
		1,
		2,
		3,
		4,
		5,
	}
	data := make([]*int, len(source))

	//Wrong way
	for i, value := range source {
		data[i] = &value
	}

	fmt.Print("Wrong got: ")
	for _, datum := range data {
		fmt.Print(*datum, " ")
	}
	fmt.Println()

	//Valid way #1
	data = make([]*int, len(source))
	for i := range source {
		data[i] = &source[i]
	}

	fmt.Print("Valid got: ")
	for _, datum := range data {
		fmt.Print(*datum, " ")
	}
	fmt.Println()
}
