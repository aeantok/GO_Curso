package main

import "fmt"

func main() {
	c := 0

	//Funciona como "while"
	for c <= 5 {
		fmt.Println(c)
		c++
	}

	fmt.Println("-----------------------------")

	//Funciona como "for"
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}
