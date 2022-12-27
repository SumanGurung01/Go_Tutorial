/*
	Date : Tue Dec 27 2022 10:00:26 GMT+0530 (India Standard Time)
	Author : Suman Gurung
	Description : Constants and %T format specifier on Go
*/

package main

import "fmt"

func main() {

	const LENGTH float64 = 10
	const BREADTH float64 = 20

	// LENGTH = 45  -> error

	area := LENGTH * BREADTH

	fmt.Println("Area is :", area)
	fmt.Printf("Type of area is %T", area)
}

/*

Constants are declare using const keyword
%T is used to print the type of variable
:= is used for dynamic type allocation

*/
