package main

import "fmt"

func main() {

	var fname, lname, nid, n1, n2, res = form("Ahmad", "Karimi", 22332, 100, 50)

	fmt.Println("Hello ", fname, lname, ". Your NID is ", nid, "You entered two numbers as ", n1, "and ", n2, "for addition, and the result of addition is ", res)

}

func form(firstName string, lastName string, nid int, num1 int, num2 int) (string, string, int, int, int, int) {
	var addResult = num1 + num2
	return firstName, lastName, nid, num1, num2, addResult
}
