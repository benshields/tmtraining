package main

import "fmt"

/* What's the value of this expression: (true && false) || (false && true) || !(false && false)? */
func main() {
	// (true && false) || (false && true) || !(false && false)
	//	false 			|| 	false 			|| !(false)
	// 	false 			||	false 			|| true
	// 	true
	fmt.Println((true && false) || (false && true) || !(false && false))
}
