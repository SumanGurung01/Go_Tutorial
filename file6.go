/*
	Date : Wed Dec 28 2022 10:09:32 GMT+0530 (India Standard Time)
	Author : Suman Gurung
	Description : Loops in Go
*/

package main

import "fmt"

func main() {

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println("i :", i)
	}

	// while loop
	age := 10
	for age < 18 {
		fmt.Println("age :", age)
		age++
	}

	// do while
	marks := 5
	for {
		fmt.Println("marks :", marks)
		marks++

		if marks > 10 {
			break
		}
	}

	var array [10]int

	for i := 0; i < 10; i++ {
		fmt.Scanln(&array[i])
	}

	fmt.Println("array :", array)
}

/*
	SYNTAX :

	1. FOR LOOP
	for initialization ; condition ; updatation {
		body ...
	}


	2. WHILE LOOP
	initialization
	for condition {
		body ...
		updatation
	}


	3. DO WHILE LOOP
	initialization
	for {
		body...
		updatation
		if condition{
			break
		}
	}

	Note : Go has only for which perform all kinds of looping
*/
