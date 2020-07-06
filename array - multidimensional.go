package main

import "fmt"

func main() {

	var arr [4][4]int
	arr[0] = [4]int{1, 1, 1, 1}
	arr[1] = [4]int{2, 2, 2, 2}
	arr[2] = [4]int{3, 3, 3, 3}
	arr[3] = [4]int{4, 4, 4, 4}

	fmt.Println("", arr[0], "\n", arr[1], "\n", arr[2], "\n", arr[3])

}
