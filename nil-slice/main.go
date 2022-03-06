package main

import "fmt"

func acceptStringSlice(data []string) {
	fmt.Printf("Got: %+v\n", data)
	for _, val := range data {
		fmt.Println(val)
	}
}

func main() {
	var nilSlice []string
	emptySlice := make([]string, 0)

	acceptStringSlice(emptySlice)
	acceptStringSlice(nilSlice)

	fmt.Println("Done")
}
