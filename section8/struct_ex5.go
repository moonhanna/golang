//구초제 심화(5)

package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executives struct {
	Employee     //is a 관계, 상속개념
	specialBonus float64
}

func main() {
	//구조체 임베디드 메소드 오버라이딩 패턴

	//예제1
	//직원
	ep1 := Employee{"kim", 2000000, 150000}
	ep2 := Employee{"park", 1500000, 200000}

	//임원
	ex := Executives{
		Employee{"lee", 5000000, 1000000},
		1000000,
	}

	fmt.Println(int(ep1.Calculate()))
	fmt.Println(int(ep2.Calculate()))

	//Employee 구조체 통해서 메소드 호출
	fmt.Println(int(ex.Calculate()))                            //메소드 오버라이딩
	fmt.Println(int(ex.Employee.Calculate() + ex.specialBonus)) //오버라이딩을 안쓴다면...

}
