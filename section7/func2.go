//함수 기초(2)

package main

import (
	"fmt"
)

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) { //매개변수의 타입이 같을 시 생략가능
	fmt.Println(a + b)
}

func multiValue(i int) {
	i *= 10
}

func multiReference(i *int) {
	*i *= 10
}

func main() {

	//함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)
	//예제1
	sum(100, add) //함수 전달

	//예제2
	//call by value
	a := 100
	multiValue(a)
	fmt.Println(a)

	//예제3
	//call by reference
	b := 100
	multiReference(&b)
	fmt.Println(b)
}
