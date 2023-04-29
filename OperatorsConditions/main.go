package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//curso()
	tarea()

}

func tarea() {
	var age int
	stdin := bufio.NewReader(os.Stdin)
	stdin.ReadString('\n')
	fmt.Print("-> ")

	_, error := fmt.Scanf("%d", &age)

	if error == nil {
		if age > 15 {
			fmt.Println("Puede seguir avanzando")
		} else {
			fmt.Println("No puede seguir circulando")
		}

	} else {
		fmt.Println(error)
	}
}

func curso() {
	fmt.Println("Curso")
	fmt.Println("Operators")

	yearsOld := 41

	fmt.Println(yearsOld > 30)
	fmt.Println(yearsOld < 30)

	fmt.Println()
	fmt.Println("OR ||")
	fmt.Println(yearsOld > 30 || yearsOld < 30)
	fmt.Println("AND &&")
	fmt.Println(yearsOld > 30 && yearsOld < 30)

	fmt.Println("NOT !")
	fmt.Println(true)
	fmt.Println(!true)
	fmt.Println(yearsOld > 30)
	fmt.Println(!(yearsOld > 30))

	fmt.Println("Conditions")

	fmt.Println("IF")
	if yearsOld > 30 {
		fmt.Printf("%d is higher than 30\n", yearsOld)
	}

	if value := true; value == true {
		fmt.Println("is true")
	}

	//reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	//text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	//text = strings.Replace(text, "\n", "", -1)
	var number int
	_, error := fmt.Scanf("%d", &number)

	fmt.Println(error)

	fmt.Println(number)

	switch number {
	case 1:
		fmt.Println(number)
	case 2:
		fmt.Println(number)
	default:
		fmt.Printf("El numero ingresado es : %d", number)
	}
}
