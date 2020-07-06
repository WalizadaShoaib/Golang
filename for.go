package main

import "fmt"

func main() {
	fmt.Println("Welcome! This program adds two number and prints the result")
	add()
}

func add() {

	for i := 1; i <= 3; i++ {

		var num1 float64
		var num2 float64
		fmt.Print("Enter the first number: ")
		fmt.Scanln(&num1)
		fmt.Print("Enter the second number: ")
		fmt.Scanln(&num2)
		var result = num1 + num2
		fmt.Print("\n***The result is*** ", result, "\n\n")

	}

}
