package main

import "fmt"

func acceptMap(data map[string]string) {
	fmt.Printf("Got: %+v\n", data)
	for _, val := range data {
		fmt.Println(val)
	}
}

func main() {
	var nilSlice map[string]string
	emptySlice := make(map[string]string, 0)

	acceptMap(emptySlice)
	acceptMap(nilSlice)

	fmt.Println("Done")
}
