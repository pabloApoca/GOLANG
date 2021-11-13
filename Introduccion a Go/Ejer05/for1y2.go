package main

import "fmt"

func main() {
	//i := 1

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		//i++
	}

	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			fmt.Println("Hacertaste, era el numero", numero)
			break
		}
	}
}
