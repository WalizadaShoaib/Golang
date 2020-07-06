package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {

		porcess()
	}

}

func porcess() {
	var firstName, LastName string
	var age int
	bioData(firstName, LastName, age)
	fmt.Print("Enter your first name:		")
	fmt.Scanln(&firstName)
	fmt.Print("Enter your last nam:		")
	fmt.Scanln(&LastName)
	fmt.Print("How old are you?		")
	fmt.Scanln(&age)
	fmt.Print("\nGreat!! Now in this phase we ask you to enter two numbers where we add the numbers for you and produce the result:\n\nEnter the first number:		")
	var n1, n2 int
	fmt.Scanln(&n1)
	fmt.Print("Enter the second number:	")
	fmt.Scanln(&n2)
	var result = add(n1, n2)
	println("\nThis is the complete data after processing.\n\nYour Full Name is  ", firstName, " ", LastName, "and you are ", age, " years old.\nYou entered two numbers as ", n1, " and ", n2, ".We added these numbers for you and the result is", result, "\n\n")

}

func bioData(firstName string, lastName string, age int) (string, string, int) {
	return firstName, lastName, age
}
func add(num1, num2 int) int {
	return num1 + num2
}

/*

func bio() {
	var name, surname string
	var age int
	fmt.Print("Enter your Name:		")
	fmt.Scanln(&name)
	fmt.Print("Enter your Surname:		")
	fmt.Scanln(&surname)
	fmt.Print("Enter your age:		")
	fmt.Scanln(&age)
	fmt.Print("\nHello ", name, " ", surname, " Your age is ", age, "\n\n")
}

*/
