package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		/*
			if i == 5 {
				fmt.Println("Se rompe el bucle")
				break
			}
		*/

		if i == 5 {
			//fmt.Println("Se rompe el bucle")
			continue
		}

		fmt.Println(i)
	}
}
