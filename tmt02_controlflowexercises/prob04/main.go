package main

import "fmt"

func main() {
	fmt.Print("Please enter a small number and a larger number: ")
	var small, larger int
	if _, err := fmt.Scan(&small, &larger); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The remainder of %v / %v is: %v\n", larger, small, larger%small)
}
