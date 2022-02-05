package main

import "fmt"

func main() {
	// *** ARRAYS ***
	var array1 [2]string
	array1[0] = "dato1"
	array1[1] = "dato2"

	fmt.Println("*** ARRAYS ***")
	fmt.Println(array1)

	array2 := [3]int{1, 2, 3}
	fmt.Println(array2)

	// *** SLICEN ***
	var slicen1 []string
	slicen1 = append(slicen1, "dato01", "dato02", "dato03", "dato04")
	fmt.Println()
	fmt.Println("*** SLICENS ***")
	fmt.Println(slicen1)
}
