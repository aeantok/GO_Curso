package main

import "fmt"

func main() {
	var (
		n      int
		salida string
	)

	fmt.Print("Ingrese un número del 1 al 5:")
	fmt.Scanln(&n)

	switch n {
	case 1:
		salida = "UNO"
	case 2:
		salida = "DOS"
	case 3:
		salida = "TRES"
	case 4:
		salida = "CUATRO"
	case 5:
		salida = "CINCO"
	default:
		salida = "Valor ingresado fuera de rango"
	}

	fmt.Println(salida)
}
