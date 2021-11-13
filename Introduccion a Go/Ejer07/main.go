package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	//Sumamos
	fmt.Printf("Sumo 5 + 7 = %d\n", Calculo(5, 7))

	//Restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 4 = %d\n", Calculo(6, 4))

	//Dividimos
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Dividimos 12 / 4 = %d\n", Calculo(12, 4))

	Operaciones()

	//CLOSURES
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(tablaDel, " * ", i, " = ", MiTabla())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 12
		return a + b
	}
	fmt.Println(resultado())
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
