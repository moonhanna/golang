//자료형 : 맵(4)

package main

import "fmt"

func main() {

	//맵(Map)
	//맵 조회 할 경우에 주의 할 점

	//예제1
	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value4, ok2 := map1["lemon"]
	value2 := map1["kiwi"] //없을 경우 기본 값
	value3, ok := map1["kiwi"]

	//두 번째 리턴 값으로 키 존재 유무 확인
	fmt.Println(value1)      //진짜 값
	fmt.Println(value4, ok2) //value가 0이여서 true 리턴
	fmt.Println(value2)      //기본 값
	fmt.Println(value3, ok)  //value가 없어서 false 리턴

	//예제2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("kiwi is not exist!")
	}

	if value, ok := map1["banana"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("banana is not exist!")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("kiwi is not exist!")
	}

}
