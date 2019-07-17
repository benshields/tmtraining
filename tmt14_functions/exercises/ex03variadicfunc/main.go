package main

import "fmt"

/* Write a function with one variadic parameter that finds the greatest number in a list of numbers. */
/* returns the max value from a slice of int, and whether the result is valid (based on legnth) */
func greatest(vs ...int) (int, bool) {
	if len(vs) <= 0 {
		return 0, false
	}
	max := vs[0]
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max, true
}

func main() {
	inp := []int{22233, 6, 3, 1, 3, 76, 8, 7, 5, 3, 2, 4, 5, 6, 5, 3, 2, 22233, 3}
	fmt.Println(greatest(inp...))
}
