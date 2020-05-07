//데이터 타입 : 문자열(1)
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//문자열
	//큰 따옴표 "", 백스쿼트 ``
	//golang : 문자 char 타입 존재 하지 않음 -> rune(int32)로 문자 코드 값으로 표현
	//문자 : ''로 작성
	//자주 사용하는 escape : \\ , \' , \" , \a(콜솔벨) , \b(백스페이스) , \f(쪽바꿈) , \n(줄바꿈) , \r(복귀) , \t(탭) ...

	//예제1
	var str1 string = "c:\\Users\\" // -> c:/users/
	str2 := `c:\Users\`             //가독성을 위해 백스쿼트 사용 가능

	fmt.Println(str1)
	fmt.Println(str2)

	//예제2
	//길이(바이트 수)
	fmt.Println(len("안녕하세요")) //3 x 5 = 15
	fmt.Println(len("Hello")) //1 x 5 = 5

	//예제3
	//실제길이
	fmt.Println(utf8.RuneCountInString("안녕하세요")) //선호 방법
	fmt.Println(utf8.RuneCountInString("Hello"))
	fmt.Println(len([]rune("안녕하세요"))) //len으로 실제 길이 구하는 법

}
