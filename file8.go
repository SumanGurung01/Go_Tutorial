/*
	Date : Thu Dec 29 2022 23:02:19 GMT+0530 (India Standard Time)
	Author : Suman Gurung
	Description : Pointers in Go
*/

package main

import "fmt"

func main() {

	var num int = 10
	var ptr *int = &num

	fmt.Println("The Address of num is", &num)
	fmt.Println("The value of ptr is", ptr)
	fmt.Println("The value stored in ptr is", *ptr)

	var array = [5]int{10, 20, 30, 40, 50}
	var arrayPtr *int = &array[0]

	fmt.Println("The value in array[0] is", arrayPtr)
	fmt.Println("The value in arrayPtr is", arrayPtr)
	fmt.Println("The value stored in arrayPtr is", *arrayPtr)

	// use array of pointers as ptr++ is not correct
	var arrayOfPtr [5]*int

	for i := 0; i < 5; i++ {
		arrayOfPtr[i] = &array[i]
	}

	for i := 0; i < 5; i++ {
		fmt.Println("The value stored in arrayOfPtr is", *arrayOfPtr[i])
	}

	// nil pointer in Go
	var p *int
	fmt.Printf("The value of p is : %x\n", p)

}

/*
note : you cannot increament pointer in Go link in C

I/O
The Address of num is 0xc000016088
The value of ptr is 0xc000016088
The value stored in ptr is 10

The value in array[0] is 0xc00000e450
The value in arrayPtr is 0xc00000e450
The value stored in arrayPtr is 10

The value stored in arrayOfPtr is 10
The value stored in arrayOfPtr is 20
The value stored in arrayOfPtr is 30
The value stored in arrayOfPtr is 40
The value stored in arrayOfPtr is 50

The value of p is : 0
*/
