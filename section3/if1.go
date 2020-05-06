//if문(1)
package main

import "fmt"

func main() {
	//제어문(조건문)
	//if문 : 반드시 boolean으로 검사 -> 1, 0 (사용불가 : 자동 형 변환 불가)
	//소괄호 사용 X

	var a = 20
	b := 20

	//예제1
	if a >= 15 {
		fmt.Println("15이상")
	}

	//예제2
	if b >= 25 {
		fmt.Println("25이상")
	}

	//예제3
	if c := true; c {
		fmt.Println("true")
	}

	//예제4
	if c := 30; c >= 20 {
		fmt.Println("20 이상")
	}

}
