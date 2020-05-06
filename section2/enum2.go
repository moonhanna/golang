//열거형2
package main

import "fmt"

func main() {
	const (
		A = iota * 10
		B
		C
	)
	fmt.Println(A, B, C)

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
	)
	fmt.Println(Jan, Feb, Mar, Apr, May, Jun)
}
