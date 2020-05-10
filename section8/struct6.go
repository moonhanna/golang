//구조체 기본(6)

package main

import (
	"fmt"
)

type Car struct { //첫 글자 대문자 : 외부에서 참조 가능
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct { //첫 글자 소문자 : 외부에서 참조 불가
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {

	//중첩 구조체
	//예제1
	car1 := Car{
		"520d",
		"silver",
		"bmw",
		spec{4000, 1000, 2000},
	}

	//내부 spec 구조체 값 출력
	fmt.Println(car1.name)
	fmt.Println(car1.color)
	fmt.Println(car1.company)
	fmt.Println(car1.detail)
	fmt.Printf("%#v\n", car1.detail)

	//예제2
	//내부 spec 구조체 필드 값 출력
	fmt.Println(car1.detail.length)
	fmt.Println(car1.detail.height)
	fmt.Println(car1.detail.width)

}
