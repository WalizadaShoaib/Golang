package main

import (
	"fmt"
)

//Calculating salary using a function

func main() {

	var s = salary{70000, 20000, 10000, 8000}
	res := cal(s)
	fmt.Println(res)

}

type salary struct {
	baseAmount         int
	bonus              int
	transportAllowance int
	tax                int
}

func cal(sal salary) int {
	var toBePaid = (sal.baseAmount + sal.transportAllowance + sal.bonus) - (sal.tax)
	return toBePaid
}

// Doing the same using a method

/*
func main() {

	var s = salary{70000, 20000, 10000, 8000}
	ress := s.calc()
	fmt.Println(ress)

}

type salary struct {
	baseAmount         int
	bonus              int
	transportAllowance int
	tax                int
}

func (sal salary) calc() int {
	var toBePaid = (sal.baseAmount + sal.transportAllowance + sal.bonus) - (sal.tax)
	return toBePaid
}
*/

// An example with pointers

/*

func main() {

	sal := salary{100000, 5000}
	res := sal.calc()
	fmt.Println(res)

}

type salary struct {
	baseAmount float64
	bonus      float64
}

func (s *salary) calc() float64 {

	amountPayable := ((*s).baseAmount) + ((*s).bonus)
	return amountPayable

}


*/

// On non-struct types - E.g 1

/*

func main() {

	n := num(3)
	res := n.calc()
	fmt.Println(res)

}

type num int

func (n num) calc() num {

	res := n * 10
	return res
}


*/

// On non-struct types - E.g 2

/*

func main() {

	n := num(3)
	res := n.calc(2, 3, 5)
	fmt.Println(res)

}

type num int

func (n num) calc(n2 num, n3 num, n4 num) num {

	res := (n * 10) + (n2 + n3 + n4)
	return res
}

*/
