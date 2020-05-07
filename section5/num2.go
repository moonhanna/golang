package main

//데이터 타입 : Numeric(2)

import "fmt"

func main() {
	//실수(부동소수점)
	//float32(7자리), float64(15자리)

	//예제1
	var num1 float32 = 0.14
	var num2 float32 = .75647
	var num3 float32 = 442.037323
	var num4 float32 = 10.0

	//지수 표기법
	var num5 float32 = 14e6
	var num6 float64 = .156875E+3
	var num7 float64 = 5.32521e-10

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(num7)

	//부동소수점 오차
	fmt.Println("------")
	fmt.Println(num4 - 0.1)
	fmt.Println(float64(num4 - 0.1)) //주의!!!

}
