//함수 심화(1)

package main

import "fmt"

func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}
	return tot
}

func sum(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}
	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println(value)
	}
}

func main() {
	//함수 고급
	//가변 인자 실습(매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음)
	//예제1
	x := multiply(5, 6, 7, 8, 9, 10)
	y := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(x)
	fmt.Println(y)

	//예제2
	prtWord("a", "apple", "test", "golang", "seoul")

	//예제3
	a := []int{1, 2, 3, 4, 5}
	m := multiply(a...) //a -> x
	n := sum(a...)
	fmt.Println(m)
	fmt.Println(n)
}
