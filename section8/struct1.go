//구조체 기본(1)

package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculator() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {

	//구조체
	//Go필드들의 집합체 또는 컨테이너
	//필드는 갖지만 메소드는 갖지 않는다
	//객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	//상속, 객체, 클래스 개념 없음
	//구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	//예제1
	kim := Account{number: "245-901", balance: 10000000, interest: 0.1}
	lee := Account{number: "245-902", balance: 12000000} //interest : 기본값 초기화
	park := Account{number: "245-903", interest: 0.2}    //balance : 기본값 초기화
	moon := Account{"245-903", 15000000, 0.3}

	fmt.Println("kim : ", kim)
	fmt.Println("lee : ", lee)
	fmt.Println("park : ", park)
	fmt.Println("moon : ", moon)

	//예제2
	fmt.Println(int(kim.Calculator()))
}
