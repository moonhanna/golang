package main

//데이터 타입 : Numeric(1)

import "fmt"

func main() {
	//데이터 타입 : 숫자형
	//정수, 실수, 복소수
	//32bit, 64bit, unsigned(양수)
	//정수 : 8진수(0), 16진수(0x), 10진수()

	var num1 = 17
	var num2 = -68
	var num3 = 0631
	var num4 = 0x32fa2c75

	fmt.Println(num1, num2, num3, num4)

	//데이터타입 : 숫자형
	//정수형 문자 출력

	//예제1
	//아스키(영문)
	var char1 byte = 72
	var char2 byte = 0110
	var char3 byte = 0x48

	fmt.Printf("%c %c %c\n", char1, char2, char3)
	fmt.Printf("%d %d %d\n", char1, char2, char3)
	fmt.Printf("%d %o %x\n", char1, char2, char3)

	//예제2
	//유니코드(한글)
	var char4 rune = 50556
	var char5 rune = 0142574
	var char6 rune = 0xC57C

	fmt.Printf("%c %c %c\n", char4, char5, char6)
	fmt.Printf("%d %d %d\n", char4, char5, char6)
	fmt.Printf("%d %o %x\n", char4, char5, char6)

}
