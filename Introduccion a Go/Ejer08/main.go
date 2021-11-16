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

	matriz[3][0] = 10
	matriz[3][2] = 20
	matriz[3][4] = 30
	matriz[3][6] = 40
	fmt.Println(matriz)

	fmt.Println("Slices: son como vectores dinamicos")
	//Donde yo puedo ampliar las dimensiones en tiempo real, si es que me quedo sin espacio
	matriz2 := []int{2, 5, 8}
	fmt.Println(matriz2)
	variante2()
	variante3()
	variante4()

}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[1:4] //Me trae los datos entre la posicion 1 y 4
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20) //Creo un slide de tipo entero que va a iniciar con un vector de largo de 5 elementos, pero va a soportar hasta 20 elementos
	fmt.Printf("Largo %d, Capacida %d", len(elementos), cap(elementos))

}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
