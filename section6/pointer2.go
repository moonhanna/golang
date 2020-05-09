//자료형 : 포인터(2)

package main

import "fmt"

func main() {

	//예제1
	i := 7
	p := &i
	fmt.Println(i, *p, &i, p, &p)

	*p++ //1증가
	fmt.Println(i, *p, &i, p, &p)

	*p = 7777 //포인터 변수 역 참조 값 변경
	fmt.Println(i, *p, &i, p, &p)

	i = 77 //원본 값 변경
	fmt.Println(i, *p, &i, p, &p)

}
