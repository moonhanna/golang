//자료형 : 배열(3)

package main

import "fmt"

func main() {

	//배열 복사
	//값 복사 확인 중요

	//예제1
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	arr2 := arr1

	fmt.Println(arr1, &arr1)
	fmt.Println(arr2, &arr2)

	fmt.Printf("%p %v\n", &arr1, arr1)
	fmt.Printf("%p %v\n", &arr2, arr2)

}
