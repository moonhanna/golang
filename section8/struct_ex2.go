//구초제 심화(2)

package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	//예제1
	kim := Account{"245-901", 10000000, 0.015}
	lee := Account{"245-902", 20000000, 0.025}
	fmt.Println(kim)
	fmt.Println(lee)

	CalculateD(kim)
	CalculateP(&lee)

	fmt.Println(int(kim.balance))
	fmt.Println(int(lee.balance))

}
