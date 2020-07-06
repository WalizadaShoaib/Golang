package main

import "fmt"

func main() {

	var a int = 100
	var b int = 200
	fmt.Println("\n**Before Swapping**")
	fmt.Println("The value at memory address ", &a, " is ", a)
	fmt.Println("The value at memory address ", &b, " is ", b)

	swap(&a, &b)

	fmt.Println("\n**After Swapping**")
	fmt.Println("The value at memory address ", &a, " is ", a)
	fmt.Println("The value at memory address ", &b, " is ", b)

}

func swap(ptrA *int, ptrB *int) {

	var tempC int = *ptrA
	*ptrA = *ptrB
	*ptrB = tempC

}
