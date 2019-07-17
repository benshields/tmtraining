package main

import "fmt"

func filter(vs []int, property func(int) bool) []int {
	rs := []int{}
	for _, v := range vs {
		if property(v) {
			rs = append(rs, v)
		}
	}
	return rs
}

func main() {
	fmt.Println(`inp = []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}`)
	inp := []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}

	fmt.Println(`fmt.Println(filter(inp, func(v int) bool {
		return v > 0
	}))`)
	fmt.Println(filter(inp, func(v int) bool {
		return v > 0
	}))

	fmt.Println(`fmt.Println(filter(inp, func(v int) bool {
		return v%2 == 1 || v%2 == -1
	}))`)
	fmt.Println(filter(inp, func(v int) bool {
		return v%2 == 1 || v%2 == -1
	}))

	fmt.Println(`fmt.Println(filter(inp, func(v int) bool {
		return v < -3 || v > 3
	}))`)
	fmt.Println(filter(inp, func(v int) bool {
		return v < -3 || v > 3
	}))
}
