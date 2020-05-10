//구초제 심화(3)

package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalculateP(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {

	//구조체 인스턴스 값 변경 시 -> 포인터 전달, 보통의 경우 -> 값 전달
	kim := Account{"245-901", 10000000, 0.015}
	lee := Account{"245-902", 20000000, 0.025}

	kim.CalculateP(1000000)
	lee.Calculate(1000000)

	fmt.Println(int(kim.balance))
	fmt.Println(int(lee.balance))

}
