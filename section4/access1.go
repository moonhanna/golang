//패키지 접근제어
package main

import (
	"fmt"
	check "section4/lib"
	"section4/mylib"
)

func main() {

	//변수, 상수, 함수, 메서드, 구조체 등 식별자
	//대문자 : 패키지 외부에서 접근 가능
	//소문자 : 패키지 외부 접근 불가(패키지 내에서만 접근 가능)
	fmt.Println(mylib.CheckNum1(101))
	fmt.Println(mylib.CheckNum2(999))

	//별칭 사용
	//빈 식별자 사용
	fmt.Println(check.CheckNum(2))
}
