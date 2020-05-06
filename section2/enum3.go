//열거형3
package main

import "fmt"

func main() {
	// _ 생략의 의미
	const (
		_ = iota
		A
		_
		C
	)

	fmt.Println(A, C)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)
	fmt.Println(DEFAULT, SILVER, GOLD, PLATINUM)
}
