package main

import "fmt"

func main() {
	var foo interface{}
	//foo = 'b'
	//foo = "boo!"
	//foo = 2
	foo = true
	switch foo.(type) {
	case int:
		fmt.Println("int")
	case string, rune:
		fmt.Println("string or rune")
	default:
		fmt.Println("type not found")
	}
}
