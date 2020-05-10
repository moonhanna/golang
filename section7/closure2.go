//함수 Closure(2)

package main

import "fmt"

func increaseCnt() func() int {
	n := 0 //지역변수(캡쳐됨)
	return func() int {
		n += 1
		return n
	}
}

func main() {

	//예제1
	cnt := increaseCnt()
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())

}
