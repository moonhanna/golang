package main

//데이터 타입 : Bool(1)

import "fmt"

func main() {
	//Bool(Boolean) : 참과 거짓
	//조건부 논리 연산자랑 주로 사용 : !, ||, &&
	//암묵적 형변환X : 0, Nil -> False 변환 없음

	//예제1
	var b1 = true
	var b2 = false
	b3 := true

	fmt.Println(b1, b2, b3)
	fmt.Println(b3 == b3)
	fmt.Println(b1 && b2)
	fmt.Println(b1 || b3)
	fmt.Println(!b3)
}
