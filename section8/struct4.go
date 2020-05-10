//구조체 기본(4)

package main

import "fmt"

func main() {

	//구조체 익명 선언 및 활용
	//예제1 type 구조체명 타입
	car1 := struct{ name, color string }{"520d", "red"}
	fmt.Println(car1)
	fmt.Printf("%#v\n", car1)

	//예제2
	cars := []struct{ name, color string }{{"528d", "red"}, {"420d", "white"}, {"380d", "blue"}}

	for _, c := range cars {
		fmt.Printf("(%s, %s) ----- (%#v)\n", c.name, c.color, c)
	}

}
