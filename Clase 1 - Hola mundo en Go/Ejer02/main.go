package main

import (
	"fmt"
	"strconv"
)

var numero int
var texto string
var status bool

func main() {
	var numero2, numero3, numero4 int

	numero2, numero3, numero4, texto, status := 2, 5, 7, "Este es mi texto", true

	var moneda int64 = 88

	numero2 = int(moneda)

	texto = fmt.Sprintf("%d", moneda)

	texto = strconv.Itoa(int(moneda))

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)

	MostrartStaus()
}

func MostrartStaus() {
	fmt.Println(status)
}
