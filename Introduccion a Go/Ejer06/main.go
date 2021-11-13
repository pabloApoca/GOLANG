package main

import "fmt"

func main() {

	/*fmt.Println(uno(5))
	numero, estado := dos(2)

	fmt.Println(numero)
	fmt.Println(estado)*/

	fmt.Println(calculo(5, 4))
	fmt.Println(calculo(5, 4, 10))

	fmt.Println(calculo2(5, 4, 10, 1))
}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	}
	return 10, true
}

func calculo(numero ...int) int {
	total := 0
	for _, num := range numero {
		total = total + num
	}
	return total
}

func calculo2(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Println("\n Item", item, " = ", num)
	}
	return total
}
