/*
	Date : Mon Dec 26 2022 20:34:47 GMT+0530 (India Standard Time)
	Author : Suman Gurung
	Description : Datatypes in Go 
*/

package main

import ("fmt" 
	"unsafe" 
	"reflect")

func main(){

	var name string = "Suman" 
	fmt.Println("The name is",name)
	fmt.Print("The name is ",name,"\n")
	fmt.Printf("The name is %s\n",name)
	fmt.Print("Size : ",unsafe.Sizeof(name)," Type : ",reflect.TypeOf(name),"\n")

	var int_num int64 = 10
	fmt.Println("Value is",int_num)
	fmt.Print("Value is ",int_num,"\n")
	fmt.Printf("Value is %v\n",int_num)
	fmt.Print("Size : ",unsafe.Sizeof(int_num)," Type : ",reflect.TypeOf(int_num),"\n")

	var float_num float64 = 20.00
	fmt.Println("Value is",float_num)
	fmt.Print("Value is ",float_num,"\n")
	fmt.Printf("Value is %f\n",float_num)
	fmt.Print("Size : ",unsafe.Sizeof(float_num)," Type : ",reflect.TypeOf(float_num),"\n")

	var char_value rune = 'a'
	fmt.Println("Value is",char_value)
	fmt.Print("Value is ",char_value,"\n")
	fmt.Printf("Value is %c\n",char_value)
	fmt.Print("Size : ",unsafe.Sizeof(char_value)," Type : ",reflect.TypeOf(char_value),"\n")
	
	// automatic type assigned
	lang := "Go" 

	fmt.Println("Value is",lang)
	fmt.Print("Value is ",lang,"\n")
	fmt.Printf("Value is %s\n",lang)
	fmt.Print("Size : ",unsafe.Sizeof(lang)," Type : ",reflect.TypeOf(lang),"\n")
}

/*

Interger Tyoe : uint8 , uint16 , uint32 , uint64 , int8 , int16 , int32 , int64 
Float Type : float32  , float64
Complex Type : complex64 , complex128
String
Go does not have char it uses rune , byte 

Note : 
1. when printing values using Println or Print no need to use %
2. but this does not print the value rune or byte so we use Printf and we need mention the format specifier

%d -> integer
%v -> values
%f -> float 
%c -> character when declared as byte or rune
%s -> string



I/O

The name is Suman
The name is Suman
The name is Suman
Size : 16 Type : string

Value is 10
Value is 10
Value is 10
Size : 8 Type : int64

Value is 20
Value is 20
Value is 20.000000
Size : 8 Type : float64

Value is 97
Value is 97
Value is a
Size : 4 Type : int32

Value is Go
Value is Go
Value is Go
Size : 16 Type : string

*/