//if문(3)
package main

import "fmt"

func main() {

	i := 100

	//if - else if - else 예제
	if i >= 120 {
		fmt.Println("120이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("100이상 120미만")
	} else if i < 100 {
		fmt.Println("100미만")
	} else {
		fmt.Println("그 외")
	}

}
