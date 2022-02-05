package main

import "fmt"

func main() {
	var a, b int

	//Entrada
	fmt.Print("Valor 1 --> ")
	fmt.Scanln(&a)

	fmt.Print("Valor 2 --> ")
	fmt.Scanln(&b)

	//Salida
	fmt.Println("a = b  ---> ", a == b)
	fmt.Println("a <> b ---> ", a != b)
	fmt.Println("a > b  ---> ", a > b)
	fmt.Println("a < b  ---> ", a < b)
	fmt.Println("a >= b ---> ", a >= b)
	fmt.Println("a <= b ---> ", a <= b)

}
