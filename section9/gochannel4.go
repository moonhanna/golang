//채널(Channel) 기초 4

package main

import (
	"fmt"
	"runtime"
)

func main() {
	//채널(Channel)
	//예제1(비동기 : 버퍼 사용)
	//버퍼 : 송신 -> 가득차면 대기, 비어있으면 작동 //가득찰 때 까지 송신
	//      수신 -> 비어있으면 대기, 가득차있으면 작동 //비어질 때 까지 수신

	runtime.GOMAXPROCS(1)
	ch := make(chan bool, 2)
	cnt := 12

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true //true 보내고 수신 될 때까지 대기
			fmt.Println("Go : ", i)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}

}
