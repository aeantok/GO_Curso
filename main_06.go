package main

import "fmt"

func main() {
	var a, b int

	//Entrada
	fmt.Print("Valor 1 --> ")
	fmt.Scanln(&a)

	fmt.Print("Valor 2 --> ")
	fmt.Scanln(&b)

	//Proceso
	suma := a + b
	resta := a - b
	multi := a * b
	divi := a / b
	modulo := a % b

	fmt.Println("a + b = ", suma)
	fmt.Println("a - b = ", resta)
	fmt.Println("a * b = ", multi)
	fmt.Println("a / b = ", divi)
	fmt.Println("a % b = ", modulo)
}
