package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		process()
	}
}

func process() {
	var nat string
	var age int
	fmt.Print("\n\nWhat is your nationality?")
	fmt.Scanln(&nat)
	fmt.Print("How old are you?")
	fmt.Scanln(&age)
	eligibility(nat, age)
}

func eligibility(nationality string, age int) (string, int) {

	if nationality != "Afghan" && age < 18 {
		fmt.Println("You are under age and not an Afghan citizen, you are not eligible")
	} else if age < 18 {
		fmt.Println("You are under age")
	} else if nationality != "Afghan" {
		fmt.Println("Only Afghans can vote")
	} else if nationality == "Afghan" && age >= 18 {
		fmt.Println("You are eligible to vote!")
	}
	return nationality, age
}
