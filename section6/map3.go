//자료형 : 맵(3)

package main

import "fmt"

func main() {

	//맵(Map)
	//맵 값 변경 및 삭제

	//예제1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
		"home1":  "http://test1.com",
	}
	fmt.Println(map1)

	map1["home2"] = "http://test2.com" //추가
	fmt.Println(map1)

	map1["home2"] = "http://test2-2.com" //수정
	fmt.Println(map1)

	//예제2(삭제)
	delete(map1, "home2")
	fmt.Println(map1)

}
