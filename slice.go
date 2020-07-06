package main

import (
	"fmt"
	"reflect"
)

func main() {

	//One way of initalizing the slice
	//var fruits = []string{"Apple", "Mango", "Orange", "Mellon"}

	//another way of initilizing the slice
	var fruits = make([]string, 4)
	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[2] = "Cherry"
	fruits[3] = "Mellon"
	//adding more items in slice with append keyword, the length and capacity of the slice are incremented or decremented
	//automatically after the append operation.
	fruits = append(fruits, "Pear", "Apricot", "Watermellon", "Grapes", "Kevi")
	fmt.Println(fruits)
	fmt.Println(reflect.TypeOf(fruits).Kind())
	fmt.Println(len(fruits), cap(fruits))

	//Accessing some limited values out of a slice

	fmt.Println(fruits[3:7])

	//Itterating over slive
	fmt.Println("\n Itterating")

	for index, value := range fruits {
		fmt.Println(index, ":", value)
	}

	//calling appending
	appending()

}

// appending two slices

func appending() {
	var studentsOfclassA = []string{"Ahmad", "Mahmood", "Jamaal"}
	var studentsOfclassB = []string{"Tuba", "Samira"}

	studentsOfclassA = append(studentsOfclassA, studentsOfclassB...) //The three dots are used for appending
	fmt.Println(studentsOfclassA)

}

// Example of Two Dimensional slice

/*
package main

import "fmt"

func main() {
	twoD()
}

//Creating a two dimensional slice
func twoD() {
	var sl = [][]string{
		{"Ahmad", "John", "Suresh", "Abdullah", "Peter"},
		{"Afghanistan", "USA", "India", "UAE", "UK"},
	}
	var caps0 = cap(sl[0])

	for i := 0; i < caps0; i++ {

		fmt.Println(sl[0][i], " is from ", sl[1][i])
	}

}
*/
