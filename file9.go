/*
	Date : Thu Dec 29 2022 23:05:19 GMT+0530 (India Standard Time)
	Author : Suman Gurung
	Description : Structure in Go
*/

package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func printPerson(person Person) {
	fmt.Println("name :", person.name)
	fmt.Println("age :", person.age)
	fmt.Println("address :", person.address)
}

func printPersonWithArray(person *Person) {
	fmt.Println("name :", person.name)
	fmt.Println("age :", person.age)
	fmt.Println("address :", person.address)
}

func main() {

	var P Person

	P.name = "Suman"
	P.age = 20
	P.address = "Siliguri"

	fmt.Println("P name :", P.name)
	fmt.Println("P age :", P.age)
	fmt.Println("P address :", P.address)

	/*-------------------------------------------------------------*/

	var S Person

	S.name = "Sam"
	S.age = 20
	S.address = "Sikkim"

	printPerson(S)

	/*-------------------------------------------------------------*/

	printPersonWithArray(&P)

}

/*
Syntax:
type structurename struct{
	variablename datatype,
	...
}
*/
