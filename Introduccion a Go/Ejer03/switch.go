package main

import "fmt"

func main() {

	switch numero := 3; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)

	default:
		fmt.Println("Es mayor a 5.")
	}
}
