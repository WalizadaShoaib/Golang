package main

import "fmt"

//Nested Structs // Think of it as table relationships
//Example 5

func main() {
	var emp employees
	emp.empId = 1
	emp.empFirstName = "Ahmad"
	emp.empLastName = "Karimi"
	emp.empDept.deptId = 55
	emp.empDept.deptName = "Admin"
	fmt.Println(emp)
}

type employees struct {
	empId        int
	empFirstName string
	empLastName  string
	empDept      departments
}

type departments struct {
	deptId   int
	deptName string
}

//Creating an instance of struct with different approaches
/* Example 4
func main() {
	//Creating an instance using new keyword
	var st1 = new(students)
	var st2 = new(students)
	//Creating an instance using the litterals
	var st3 = students{3, "Sama"}

	//First declaring and then initializing
	var st4 students
	st4.SID = 4
	st4.Name = "Jamila"

	st1.SID = 1
	st1.Name = "Ahmad"
	st2.SID = 2
	st2.Name = "Asad"

	// Printing in differnt formats
	fmt.Println("Id", st1.SID)
	fmt.Println("Name", st1.Name)
	fmt.Println("Id", st2.SID)
	fmt.Println("Name", st2.Name)
	fmt.Println("Id", st3.SID, "Name", st3.Name)
	fmt.Println(st3)
	fmt.Println(st4.SID, st4.Name)

}

type students struct {
	SID  int
	Name string
}
*/
// working with pointers in sturct, its same as working with pointers in other types
/* Example 3
func main() {

	var emp1 = employees{1, "Ahmad"}
	var ptr *employees
	ptr = &emp1

	fmt.Println(ptr)
	fmt.Println(ptr.Name)
	fmt.Println(&ptr)

}

type employees struct {
	EId  int
	Name string
}
*/
/*Example 2

func main() {
	details()
}

func details() {
	var emp1 employees
	var emp2 employees
	var emp3 employees

	emp1.Id = 1
	emp1.Name = "Ahmad"
	emp1.Department = "Finance"
	emp2.Id = 2
	emp2.Name = "Sama"
	emp2.Department = "ICT"
	emp3.Id = 3
	emp3.Name = "Jawaad"
	emp3.Department = "Finance"



	fmt.Println(emp1.Id)
	fmt.Println(emp1.Name)
	fmt.Println(emp1.Department)

	fmt.Println(emp2.Id)
	fmt.Println(emp2.Name)
	fmt.Println(emp2.Department)

	fmt.Println(emp3.Id)
	fmt.Println(emp3.Name)
	fmt.Println(emp3.Department)



}

type employees struct {
	Id         int
	Name       string
	Department string
}


*/

/* Example 1
------------


func main() {
	var book1 = books{1, "C", "Ahmad"}
	var book2 = books{2, "Java", "Mahmood"}
	var book3 = books{3, "Go", "Peter"}

	details(book1)
	details(book2)
	details(book3)
}

//Passing struct to function, Its just like any other type
func details(book books) {

	fmt.Println(book.BID)
	fmt.Println(book.Title)
	fmt.Println(book.Author)

}

type books struct {
	BID    int
	Title  string
	Author string
}

*/
