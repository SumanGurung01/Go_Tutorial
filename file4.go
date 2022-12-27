/*
	Date : Tue Dec 27 2022 10:23:23 GMT+0530 (India Standard Time)
	Author : Suman Gurung
	Description : Decision Making in Go
*/

package main

import "fmt"

func main() {

	num := 10

	if num > 100 {
		fmt.Println("num > 100")
	} else if num < 100 {
		fmt.Println("num < 100")
	} else {
		fmt.Println("num = 100")
	}

	marks := 100

	switch marks {
	case 100:
		fmt.Println("A")
	case 90:
		fmt.Println("B")
	case 80:
		fmt.Println("C")
	case 70:
		fmt.Println("D")
	case 60:
		fmt.Println("E")
	default:
		fmt.Println("fefault case")
	}

	num1 := 10

	switch {
	case num1 == 10:
		fmt.Println("case 1 is true")
	case num1 == 20:
		fmt.Println("case 2 is true")
	case num1 == 30:
		fmt.Println("case 3 is true")
	default:
		fmt.Println("default is true")
	}
}

/*

switch (var)
{
	case (value of var) :
		statement
	case (value of var) :
		statement
	...
	default :
		statement
}

OR

switch
{
	case (expression) :
		statement
	case (expression) :
		statement
	...
	default :
		statement
}

*/
