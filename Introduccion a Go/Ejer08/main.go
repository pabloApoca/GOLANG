package main

import "fmt"

//Arreglo
var tabla [10]int

//vector
var matriz [5][7]int

func main() {
	fmt.Println("----Arreglos Estaticos y Slices----")

	fmt.Println("Vectores")

	//Todas la posiciones de un arreglo se inician en cero
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	//Puedo definir el valor de la posiciones de una tabla
	tabla2 := [10]int{7, 6, 1, 2, 5, 6, 8, 3, 33, 2}
	fmt.Println(tabla2)

	//Recorrer una tabla
	for i := 0; i < len(tabla2); i++ {
		//fmt.Println(tabla2[i])
	}

	fmt.Println("Matrices")

	matriz[3][5] = 1
	fmt.Println(matriz)
}
