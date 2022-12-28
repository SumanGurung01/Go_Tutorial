/*
	Date : Wed Dec 28 2022 10:09:32 GMT+0530 (India Standard Time)
	Author : Suman Gurung
	Description : Array and I/O in Go
*/

package main

import "fmt"

func main() {

	var array [10]float32

	var num float32
	fmt.Print("Enter Number: ")
	fmt.Scanln(&num)

	array[0] = num

	fmt.Println("Array : ", array)
}

/*
	Array declaration:
	1. var array [10] int;
	2. var array = [10] int {1,2,3,...,10}  [initializing]

	1. var array[10] int = {1,2,3,...,10} [incorrect]


	Array Printing
	fmt.Println(array)

	Array Indexing : array[0] , array[1] ...
*/
