package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1, numero2, resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese el primer numero: ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese el segundo numero: ")
	fmt.Scanln(&numero2)

	fmt.Println("Descripcion: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()

	}
	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)

	//fmt.Println("La suma entre los dos numeros es: ", numero1+numero2)

}
