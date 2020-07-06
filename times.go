package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		process()
	}
}
func process() {
	fmt.Println("Enter a number")
	var num int
	fmt.Scanln(&num)
	times(num)
}
func times(j int) {

	for i := 1; i <= 10; i++ {
		r := j * i
		println(j, " * ", i, " = ", r)

	}
}
