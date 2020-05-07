//데이터 타입 : Numeric 연산(1)
package main

import (
	"fmt"
	"math"
)

func main() {
	//숫자 연산(산술, 비교)
	//타입이 같아야 가능
	//다른 타입끼리는 반드시 형 변환 후 연산
	//형 변환 없을 경우 예외(에러) 발생
	// +, -, * ,%, |, <<, >>, &, ^

	//예제1
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

}
