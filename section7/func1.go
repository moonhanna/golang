//함수 기초(1)

package main

import (
	"fmt"
	"strconv"
)

//함수 선언 위치는 어느 곳이든 가능
func helloGolang() {
	fmt.Println("Hello Golang!")
}

func sayOne(m string) {
	fmt.Println(m)
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	//함수
	//선언 : func 키워드로 선언
	//func 함수명(매개변수) (반환타입 or 반환값 변수명) : 반환 값 존재
	//func 함수명() (반환타입 or 반환값 변수명) : 매개변수 없음, 반환 값 존재
	//func 함수명(매개변수) : 매개변수 존재, 반환 값 없음
	//func 함수명() : 매개변수 없음, 반환 값 없음
	//타 언어와 달리 반환값 (return value) 다수 가능

	//예제1
	helloGolang()

	//예제2
	sayOne("Hello World!")

	//예제3
	fmt.Println(sum(3, 5))
	fmt.Println(strconv.Itoa(sum(3, 5))) //숫자 -> 문자열로 변경
}
