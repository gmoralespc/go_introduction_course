package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int

	fmt.Println(myIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	var myArrayVar [6]int

	fmt.Println(myArrayVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar)*8)

	myArrayVar1 := [3]string{"value1", "value2", "value3"}

	fmt.Println(myArrayVar1)
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myArrayVar1, unsafe.Sizeof(myArrayVar1), unsafe.Sizeof(myArrayVar1)*8)

	myArrayVar[0] = 2
	myArrayVar[1] = 5
	myArrayVar[2] = 9

	fmt.Println(myArrayVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar)*8)

	fmt.Println("Size: ", len(myArrayVar))

	myArrayVar[3] = 99

	fmt.Println(myArrayVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar)*8)

	fmt.Println("Size: ", len(myArrayVar))

	//Slices
	fmt.Println("Slices")
	var slice1 []int
	fmt.Println(slice1)

	slice1 = append(slice1, 10, 20, 30, 40, 50)
	fmt.Println(slice1)

	mySlice := myArrayVar[2:4]

	fmt.Println(mySlice)
	fmt.Println("Size: ", len(mySlice))
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", mySlice, unsafe.Sizeof(mySlice), unsafe.Sizeof(mySlice)*8)
	fmt.Println(&mySlice)
	fmt.Println(&mySlice[0])
	fmt.Println(&myArrayVar[2])
	fmt.Println(&mySlice[1])
	fmt.Println(&myArrayVar[3])

	mySlice[0] = 80
	mySlice[1] = 100

	fmt.Println(mySlice)
	fmt.Println(myArrayVar)

	fmt.Println(myArrayVar[:4])
	fmt.Println(myArrayVar[2:])

	slice := make([]int, 3)

	fmt.Println(slice)
	fmt.Println("Size: ", len(slice))
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", slice, unsafe.Sizeof(slice), unsafe.Sizeof(slice)*8)
}
