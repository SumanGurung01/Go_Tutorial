/*
	Date : Wed Dec 28 2022 10:09:32 GMT+0530 (India Standard Time)
	Author : Suman Gurung
	Description : Array and I/O in Go
*/

package main

import "fmt"

func main() {

	var array [10]float32
	newArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var anotherArray [10]int

	var num float32
	fmt.Print("Enter Number: ")
	fmt.Scanln(&num)

	array[0] = num

	fmt.Println("array : ", array)
	fmt.Println("newArray : ", newArray)
	fmt.Println("anotherArray : ", anotherArray)
}

/*

	I/O:
	Enter Number : 30
	array : [30 0 0 0 0 0 0 0 0 0]
	newArray : [1 2 3 4 5 6 7 8 9 10]
	anotherArray : [0 0 0 0 0 0 0 0 0 0]


	Array declaration:
	1. var array [10] int;
	2. var array = [10] int {1,2,3,...,10}  [initializing]

	1. var array[10] int = {1,2,3,...,10} [incorrect]


	Array Printing
	fmt.Println(array)

	Array Indexing : array[0] , array[1] ...

	Note : array are initialized with 0's at declaration
*/
