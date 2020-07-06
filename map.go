package main

import (
	"fmt"
	"runtime"
)

func main() {

	emps()
	arr()
	mk()
	del()
	edata()
}

func emps() {
	var e = map[int]string{
		1: "Ahmad",
		2: "Mahmood",
		3: "Sama",
	}

	fmt.Println(e[1]) // printing a single value of map
	fmt.Println(e)    // printing all values

	elem, ok := e[1] // checking if an element exist in map
	fmt.Println(elem, ok)

}

// The difference between map, array/slices etc are in key, map has a key and you can look for values baased on
// a key, however you can look for array elements based on index, check below example of an array

func arr() {
	var e = [3]string{
		"Ahmad",
		"Mahmood",
		"Sama",
	}
	fmt.Println(e[0])
	fmt.Println(e)
}

// Declaring Map with make keyword

func mk() {
	var students = make(map[int]string)
	students[11] = "John"
	students[22] = "Peter"
	students[33] = "Shah"
	fmt.Println(len(students))            // checking the length of map
	fmt.Println(students)                 // retrieveing all elements
	fmt.Println(students[1], students[3]) // retrieving based on key

	//Itterating over a map

	fmt.Println(" \n\nItterating")

	for key, value := range students {
		fmt.Println(key, value)
	}
	fmt.Println("\nEnd of itteration")

}

//Deleting from map

func del() {
	var students = make(map[int]string)
	students[1] = "John"
	students[2] = "Peter"
	students[3] = "Shah"
	fmt.Println("\n\n\n\nBefore Delete")
	fmt.Println(students) // Befor Delete

	delete(students, 2) // Deleting by calling delete function
	fmt.Println("\n\nAfter Delete")
	fmt.Println(students)

}

//Passing struct into a map

type employees struct {
	ID         int
	Name       string
	Department string
	Contact    int
}

func edata() {
	var mp = make(map[int]employees)
	mp[1] = employees{1, "Ahmad", "ICT", 783838383}
	mp[2] = employees{2, "Sama", "Finance", 7987383}

	fmt.Println("\n\n\n", mp)

	fmt.Println(mp[1])

	fmt.Println("Processers used", runtime.NumCPU()) // This checks how many processors are used
}
