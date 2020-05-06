//for문(3)
package main

import "fmt"

func main() {

	//예제 1 -> Loop를 통해 가장 바깥으로 빠져나갈 수 있음, 그냥 break 시 안쪽 for문만 빠져나가고 바깥쪽 for문은 계속 실행됨
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println(i, j)
		}
	}

	//예제2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//예제3 //자주쓰이는 패턴은 아님...
Loop2:
	//에러 발생(Loop 레이블 밑에 관련이 없는 소스코드가 있을 시)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println(i, j)
		}
	}

}
