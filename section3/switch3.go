//switch문(3)
package main

import "fmt"

func main() {
	//예제1
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		if a >= 20 {
			fmt.Println("20보다 크다")
		}
		fmt.Println("a -> ", a, "는 홀수")
	}

	//예제2
	//fallthrough -> 반드시 실행해야 할 때 (조건문에 걸린 후 부터 실행됨)
	switch e := "go"; e {
	case "java":
		fmt.Println("java")
		fallthrough
	case "go":
		fmt.Println("go")
		fallthrough
	case "js":
		fmt.Println("js")
		fallthrough
	case "c":
		fmt.Println("c")
	}
}
