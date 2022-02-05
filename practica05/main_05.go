package main

import "fmt"

func main() {
	var (
		nombre string
		edad   int
		pi     float32
	)

	fmt.Print("Ingrese su nombre: ")
	//Pide el ingreso por pantalla.
	//Primer parámetro: %s - indica el tipo de dato a ingresar
	//Segundo parámetro: &variable - espacio donde se guarda el valor ingresado (el símbolo "&" es obligatorio)
	fmt.Scanln(&nombre)

	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&edad)

	fmt.Print("Ingrese el valor de PI: ")
	fmt.Scanln(&pi)

	fmt.Printf("------------------------------\n")
	fmt.Printf("Nombre: %s Edad: %d\nValor de Pi: %f", nombre, edad, pi)

}
