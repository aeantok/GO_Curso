package main

import (
	"fmt"
	"time"
)

func main() {
	//saludar()
	var (
		opc             int
		a, b, resultado float32
		signo           string
	)

	fmt.Print("Ingrese el valor 1: ")
	fmt.Scanln(&a)
	fmt.Print("Ingrese el valor 2: ")
	fmt.Scanln(&b)
	fmt.Println("")
	fmt.Println("---------------------------")
	fmt.Println("Elija la operación")
	fmt.Println("---------------------------")
	fmt.Println("1: Sumar")
	fmt.Println("2: Restar")
	fmt.Println("3: Multiplicar")
	fmt.Println("4: Dividir")
	fmt.Println("---------------------------")
	fmt.Print("Opción: ")
	fmt.Scanln(&opc)

	switch opc {
	case 1:
		resultado = sumar(a, b)
		signo = "+"
	case 2:
		resultado = restar(a, b)
		signo = "-"
	case 3:
		resultado = multi(a, b)
		signo = "*"
	case 4:
		resultado = dividir(a, b)
		signo = "/"
	}

	fmt.Printf("%f %s %f = %f", a, signo, b, resultado)

}

func saludar() {
	//var nombre, momento string
	var momento string
	t := time.Now()
	h := t.Local().Hour()

	if h >= 6 && h <= 12 {
		momento = "buen día!"
	} else if h > 12 && h < 20 {
		momento = "buenas tardes!"
	} else {
		momento = "buenas noches!"
	}

	//fmt.Print("Ingrese su nombre: ")
	//fmt.Scanln(&nombre)
	//fmt.Printf("Hola %s, %s", nombre, momento)
	nombre, edad := datos()
	fmt.Printf("Hola %s [%d años], %s", nombre, edad, momento)
}

//Función con retorno de datos
func datos() (string, int) {
	var (
		nombre string
		edad   int
	)

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&edad)

	return nombre, edad
}

func sumar(a float32, b float32) float32 {
	return a + b
}

func restar(a float32, b float32) float32 {
	return a - b
}

func multi(a float32, b float32) float32 {
	return a * b
}

func dividir(a float32, b float32) float32 {
	return a / b
}
