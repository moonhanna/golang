//데이터 타입 : 문자열(2)
package main

import (
	"fmt"
)

func main() {
	//문자열 표현
	//문자열 : UTF-8 인코딩 (유니코드 문자 집합) -> 바이트 수 주의

	//예제1
	var str1 = "golang"
	var str2 = "고프로그래밍"

	fmt.Println(str1[0], str2[0]) //정수값 출력
	fmt.Printf("%c %c\n", str1[0], str2[0])

	conStr := []rune(str2)
	fmt.Printf("%c\n", conStr[0]) //한글 정상 출력

	//예제2
	for i, char := range str1 {
		fmt.Printf("%c(%d) ", char, i)
	}

}
