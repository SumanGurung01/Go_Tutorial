/*
	Date : Wed Dec 28 2022 10:55:32 GMT+0530 (India Standard Time)
	Author : Suman Gurung
	Description : Functions in Go
*/

package main

import "fmt"

func max(num1 int, num2 int) {
	fmt.Print("Max : ")
	if num1 > num2 {
		fmt.Println(num1)
		return
	}
	fmt.Println(num2)
	return
}

func sum(array [10]int, n int) int {
	var result int = 0
	for i := 0; i < n; i++ {
		result = result + array[i]
	}
	return result
}

func main() {

	var num1, num2 int
	var array [10]int

	fmt.Scanln(&num1, &num2)
	max(num1, num2)

	for i := 0; i < 10; i++ {
		fmt.Scanln(&array[i])
	}

	fmt.Println("Sum of Array : ", sum(array, 10))
}

/*
SYNTAX:

func function_name ( argument_list) return_type{
	body...

	return
}

*/
