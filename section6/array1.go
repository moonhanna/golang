//자료형 : 배열(1)

package main

import "fmt"

func main() {

	//배열
	//배열은 용량, 길이 항상 같다.
	//배열 vs 슬라이스 차이점 중요
	//길이고정 vs 길이가변
	//값 타입 vs 참조 타입
	//복사 전달 vs 참조 값 전달
	//전체 비교연산자 사용 가능 vs 비교연산자 사용불가
	//대부분 슬라이스 사용한다.

	//cap() : 배열, 슬라이스 용량
	//len() : 배열, 슬라이스 개수

	//예제1
	var arr1 [5]int
	var arr2 = [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3, 4, 5} //배열크기 자동 맞춤, 이렇게 잘 안씀...
	arr7 := [5][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	fmt.Printf("%d %v\n", len(arr3), arr3)
	fmt.Printf("%d %v\n", len(arr4), arr4)
	fmt.Printf("%d %v\n", len(arr5), arr5)
	fmt.Printf("%d %v\n", len(arr6), arr6)
	fmt.Printf("%d %v\n", len(arr7), arr7)

	//예제2
	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"kim", "lee"}
	fmt.Printf("%d %v\n", len(arr8), arr8)
	fmt.Printf("%d %v\n", len(arr9), arr9)
	fmt.Printf("%d %v\n", len(arr10), arr10)
}
