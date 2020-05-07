//데이터 타입 : 문자열 연산(1)
package main

import (
	"fmt"
	"strings"
)

func main() {
	//문자열 연산
	//추출, 비교, 조합(결합)

	//예제1(추출)
	var str1 = "golang"
	var str2 = "world"

	fmt.Println(str1[0:2], str1[0])
	fmt.Println(str2[3:], str2[0])
	fmt.Println(str2[:4])

	//예제2(비교)
	str3 := "golang"
	str4 := "world"

	fmt.Println(str3 == str4)
	fmt.Println(str3 != str4)
	fmt.Println(str3 > str4)
	fmt.Println(str3 < str4) //go 문자열 -> 아스키 코드에 대한 사전식 비교

	//예제3(결합 : 일반연산)
	str5 := "The Go programming language is an open source project to make programmers more productive." +
		"Go is expressive, concise, clean, and efficient."
	str6 := "Instructions for downloading and installing the Go compilers, tools, and libraries."

	fmt.Println(str5 + str6)

	//예제4(결합 : Join) //추천
	strSet := []string{} //슬라이스 선언
	strSet = append(strSet, str5)
	strSet = append(strSet, str6)

	fmt.Println(strings.Join(strSet, "--------"))
}
