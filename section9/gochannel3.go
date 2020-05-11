//채널(Channel) 기초 3

package main

import (
	"fmt"
	"time"
)

func main() {
	//채널(Channel)
	//예제1(동기 : 버퍼 미사용)
	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true //true 보내고 수신 될 때까지 대기
			fmt.Println("Go : ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main : ", i)
	}

}
