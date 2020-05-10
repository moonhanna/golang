//구조체 기본(2)

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

	//다양한 선언 방법
	//&struct, struct : &struct 포인터를 받고 역참조를 또 하기 때문에 속도는 조금 느리다
	//인터페이스 메소드를 선언만 해둔 후 오버라이딩해서 메소드에 포인터 리시버를 사용할 경우 반드시 &struct로 넘겨야함

	//예제2
	//선언 방법1
	var kim *Account = new(Account)
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	//선언 방법2
	hong := &Account{"245-903", 15000000, 0.3}

	//선언 방법3
	lee := new(Account)
	lee.number = "245-902"
	//...

	fmt.Println(kim)
	fmt.Println(hong)
	fmt.Println(lee)
	fmt.Printf("%#v", lee) //디버깅 시 용이
	fmt.Println("\n")

	//예제2
	fmt.Println(int(kim.Calculator()))

}
