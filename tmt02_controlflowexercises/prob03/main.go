package main

import "fmt"

func main() {
	fmt.Print("Please enter your name: ")
	var name string
	if _, err := fmt.Scan(&name); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hello", name)
}
