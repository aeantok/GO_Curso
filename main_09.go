package main

import "fmt"

func main() {
	fmt.Println("***** NOT *****")
	fmt.Println(!true)

	fmt.Println("")

	fmt.Println("***** AND *****")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	fmt.Println("")

	fmt.Println("***** OR *****")
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

}
