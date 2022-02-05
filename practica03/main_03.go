package main

import "fmt"

//Constante declarada a nivel general
const g1 = "Valor General"

func main() {
	const n = 25

	var nombre string
	nombre = "Ariel"

	fmt.Println(nombre)

	var a, b int = 1, 2
	var c, d int
	c = 10
	d = 20
	fmt.Println(a, b, c, d)

	var ( //Declaración de variables de distintos tipos de datos
		pi     float64
		bol    bool
		cadena = "Texto 01"
		edad   = 55
	)

	pi = 3.141592
	bol = true

	fmt.Println(pi, bol, cadena, edad)

	//Declaración dinámica de variable con asignación de dato
	v1 := 25
	v2 := "Texto 02"

	fmt.Println(v1, v2)

}
