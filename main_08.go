package main

import "fmt"

func main() {
	var n int

	fmt.Print("Ingrese un número: ")
	fmt.Scanln(&n)

	if n == 0 {
		fmt.Println(n, "es Neutro")
	} else if n%2 == 0 {
		fmt.Println(n, "es Par")
	} else {
		fmt.Println(n, "es Impar")
	}
}
