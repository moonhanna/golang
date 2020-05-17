//고루틴 동기화 기초(5)

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//고루틴 동기화 객체
	//동기화 상태(조건) 메소드 사용
	//Wait, notify, notyfyall : 기타 언어
	//Wait, signal, broadcast : go언어

	//시스템 전제 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) //비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("고루틴 기다림", n)
			condition.Wait() //고루틴 대기(멈춤)
			fmt.Println("대기 끝", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
		//fmt.Println("받음", <-c)
	}

	/*
		for i := 0; i < 5; i++ {
			mutex.Lock()
			fmt.Println("고루틴 깨움 : ", i)
			condition.Signal() //한개씩 깨움(모든 고루틴 생성 후)
			mutex.Unlock()
		}
	*/

	mutex.Lock()
	fmt.Println("고루틴 깨움")
	condition.Broadcast() //한번에 깨움
	mutex.Unlock()

	time.Sleep(time.Second * 2)

}
