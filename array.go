package main

import "fmt"

func main() {

	var i int
	var a [10]int

	/* initialize elements of array n to 0 */
	for i = 0; i < 10; i++ {
		a[i] = i + 100
		fmt.Println("Element at a [", i, "]", " = ", a[i])
	}

}

/*

func main() {

	var emp = [5]string{"Ahmad", "Asad", "Sama", "Kabir", "Jamshid"}
	//fmt.Println(emp)

	//Foreach
	for _, e := range emp {
		fmt.Println("Employee Name", e)
	}
}

*/

/*
E.g 2
func main() {

	var fruits = []string{"Apple", "Orange", "Millon"}
	fmt.Println(fruits[0])
}
*/
