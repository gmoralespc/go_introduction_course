package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int

	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	var myUintVar uint

	myUintVar = 12

	fmt.Println(myUintVar)

	var myStringVar string

	myStringVar = "My string variable"

	fmt.Println(myStringVar)

	fmt.Println("Address is: ", &myStringVar)

	var myBoolVar bool

	myBoolVar = true
	fmt.Println(myBoolVar)

	myIntVar2 := 14
	fmt.Println(myIntVar2)

	const myConst = "a"
	fmt.Println("My const is: ", myConst)

	const myIntConst int = 1
	fmt.Println("My myIntConst is: ", myIntConst)

	var my8BitsVar int8

	fmt.Printf("Int default value: %d\n", my8BitsVar)

	fmt.Printf("type: %T, bytes: %d, bits: %d \n", my8BitsVar, unsafe.Sizeof(my8BitsVar), unsafe.Sizeof(my8BitsVar)*8)
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myIntConst, unsafe.Sizeof(myIntConst), unsafe.Sizeof(myIntConst)*8)
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myStringVar, unsafe.Sizeof(myStringVar), unsafe.Sizeof(myStringVar)*8)

	var myFloat32Var float32
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	//scope
	{
		var myVarScope int
		myVarScope = 1
		fmt.Println(myVarScope)

		fmt.Println()
		fmt.Println("1 floatStrVar -----------------------------------------------------------------------------------------------------")
		floatVar := 33.11
		fmt.Printf("type: %T, bytes: %d, bits: %d \n", floatVar, unsafe.Sizeof(floatVar), unsafe.Sizeof(floatVar)*8)
		floatStrVar := fmt.Sprintf("%f", floatVar)
		fmt.Printf("type: %T, bytes: %d, bits: %d \n", floatStrVar, unsafe.Sizeof(floatStrVar), unsafe.Sizeof(floatStrVar)*8)
		fmt.Println("2 intStrVar ------------------------------------------------------------------------------------------------------")
		intVar := 22
		fmt.Printf("type: %T, bytes: %d, bits: %d \n", intVar, unsafe.Sizeof(intVar), unsafe.Sizeof(intVar)*8)
		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, bytes: %d, bits: %d \n", intStrVar, unsafe.Sizeof(intStrVar), unsafe.Sizeof(intStrVar)*8)

		intVal1, err := strconv.ParseInt("12345332", 0, 64)
		fmt.Println(err)
		fmt.Println(intVal1)
		fmt.Printf("type: %T, bytes: %d, bits: %d \n", intVal1, unsafe.Sizeof(intVal1), unsafe.Sizeof(intVal1)*8)

		intVal2, err := strconv.ParseInt("12as3asdf4dasf5332", 0, 64)
		fmt.Println(err)
		fmt.Println(intVal2)
		fmt.Printf("type: %T, bytes: %d, bits: %d \n", intVal2, unsafe.Sizeof(intVal2), unsafe.Sizeof(intVal2)*8)

		floatVar1, _ := strconv.ParseFloat("3232.23245234", 64)
		fmt.Println(floatVar1)
		fmt.Printf("type: %T, bytes: %d, bits: %d \n", floatVar1, unsafe.Sizeof(floatVar1), unsafe.Sizeof(floatVar1)*8)
	}

}
