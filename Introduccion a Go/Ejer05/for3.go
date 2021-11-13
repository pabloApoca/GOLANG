package main

import "fmt"

func main() {
	var i = 0
	for i < 10 {
		fmt.Printf("\n Valor de i: %d", i)
		if i == 5 {
			i = i * 2
			fmt.Printf(" multiplicamos i por 2 : %d", i)
			continue
		}
		fmt.Printf("		Paso por aqui. \n")
		i++
	}
}
