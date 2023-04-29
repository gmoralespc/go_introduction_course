package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum++
		fmt.Println(sum)
	}

	sum = 1

	for sum < 1000 {
		sum++
	}

	fmt.Println(sum)

	sum = 0

	for {
		if sum > 1000 {
			break
		}

		sum++
	}

	fmt.Println(sum)

	myArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := range myArr {
		fmt.Println("Index: ", i, " - Value: ", myArr[i])
	}
}
