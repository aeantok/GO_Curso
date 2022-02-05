package main

import "fmt"

func main() {
	//Muestra por pantalla y no hace un salto de línea
	// el parámetro \t hace un espacio de tabulador
	// el parámetro \n hace un salto de línea
	fmt.Print("Hola\tAriel\n")

	nombre := "Ariel"
	edad := 55
	pi := 3.141592
	fmt.Println("Nombre:", nombre, " - Edad:", edad, "\nValor de PI:", pi)
	fmt.Println("-----------------------")

	//Printf: Permite armar un texto completo usando los parámetros de tipos de datos
	//    %s: Valor de tipo string
	//    %d: Valor de timpo int
	//    %f: Valor de tipo float
	fmt.Printf("Nombre: %s Edad: %d \nValor de PI: %f", nombre, edad, pi)

}
