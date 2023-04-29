package main

import (
	"fmt"
	"unsafe"
)

func main() {
	myMap := make(map[int]string)

	myMap[1] = "A"
	myMap[2] = "B"
	myMap[3] = "C"
	myMap[4] = "D"

	fmt.Println(myMap)
	fmt.Println("Size: ", len(myMap))
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myMap, unsafe.Sizeof(myMap), unsafe.Sizeof(myMap)*8)

	delete(myMap, 2)
	fmt.Println(myMap)
	fmt.Println("Size: ", len(myMap))
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myMap, unsafe.Sizeof(myMap), unsafe.Sizeof(myMap)*8)

	myMap[1] = "2A"

	fmt.Println(myMap)
	fmt.Println("Size: ", len(myMap))
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myMap, unsafe.Sizeof(myMap), unsafe.Sizeof(myMap)*8)

	myMap[7] = ""

	fmt.Println(myMap)
	fmt.Println("Size: ", len(myMap))
	fmt.Printf("type: %T, bytes: %d, bits: %d \n", myMap, unsafe.Sizeof(myMap), unsafe.Sizeof(myMap)*8)

	v, ok := myMap[8]

	fmt.Println("The value: ", v, "Present?", ok)

	myMap2 := map[int]string{
		4: "A",
		5: "C",
		6: "Z",
	}

	fmt.Println(myMap2)
	fmt.Println("Size: ", len(myMap2))
	delete(myMap2, 4)
	fmt.Println(myMap2)
	fmt.Println("Size: ", len(myMap2))

	myMap2[5] = "PP"
	fmt.Println(myMap2)
	fmt.Println("Size: ", len(myMap2))
}
