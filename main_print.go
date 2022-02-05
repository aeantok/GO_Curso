package main

import (
	"fmt"
	"strconv"
)

func main_Print() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}
}
