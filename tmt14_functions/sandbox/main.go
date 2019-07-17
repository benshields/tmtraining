package main

import "fmt"

func main() {
	var input = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(sum(input...)) // unpack collection to pass args to variadic func

	var seconds = 2*3600 + 17*60 + 42
	hour, min, sec := tohms(seconds)
	fmt.Printf("%v seconds ---> %vh %vm %vs\n", seconds, hour, min, sec)

	inc := safecount()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}

func sum(sn ...int) int { // variadic function
	var n int
	for _, v := range sn {
		n += v
	}
	return n
}

/* converts time from seconds to hours:minutes:seconds */
func tohms(total int) (hours, mins, seconds int) { // named multiple returns
	hours = total / 3600
	total %= 3600
	mins = total / 60
	total %= 60
	seconds = total % 60
	return
}

func safecount() func() int {
	seed := 42
	return func() int {
		seed++
		return seed
	}
}
