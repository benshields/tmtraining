package main

import "fmt"

/* returns the quotient of v/2 and whether v was even */
func half(v int) (int, bool) {
	return v / 2, v%2 == 0
}

func main() {
	inp := []int{-4, -3, -2, -1, 0, 1, 2, 3, 4}

	for _, v := range inp {
		quotient, even := half(v)
		fmt.Printf("half(%v) = %v, %v\n", v, quotient, even)
	}
}
