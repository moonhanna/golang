//자료형 : 포인터(1)

package main

import "fmt"

func main() {

	//포인터
	//Go : 포인터 지원(c)
	//변수의 지역성, 연속된 메모리 참조 ..., 힙, 스택
	//파이썬, 자바 (JRE) -> 컴파일러, 인터프리터
	//포인터지원x (파이썬, c#, Java 등)
	//주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그 방지)
	//*(에스터리스크)로 사용
	//nil 로 초기화(nil == 0)

	//예제1
	var a *int            //방법1
	var b *int = new(int) //방법2

	fmt.Println(a) //nil
	fmt.Println(b) //주소값

	i := 7
	fmt.Println(i, &i) //값과 주소값 출력

	a = &i //주소값 전달
	b = &i
	fmt.Println(a, &i, &a, *a) //*a -> 역참조
	fmt.Println(b, &i, &b, *b)

}
