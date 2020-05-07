//데이터 타입 : Numeric 연산(2)
package main

import (
	"fmt"
	"math"
)

func main() {

	//예제1
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println(n1 << 2)
	fmt.Println(n1 >> 2)
	fmt.Println(^n2)

	//예제2(오버플로우 에러 : 범위 초과)
	var n1 uint8 = math.MaxUint8 + 1

}
