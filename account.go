package main

import "fmt"

// Elaborate this example and develop a proper bank account system

func main() {
	acc()
}

func acc() {

	var amount = 1000
	var withdraw int
	for i := amount; i >= 0; i-- {
		if amount > 0 {
			fmt.Print("Your account balance is ", amount, " How much would you like to withdraw?")
			fmt.Scanln(&withdraw)
			var available = amount - withdraw
			fmt.Print("Thank you for the transaction, You withdrawn ", withdraw, " and your available balance is ", available)
			amount = amount - withdraw
		} else {
			fmt.Print("Not Sufficient Balance")
		}

	}
}
