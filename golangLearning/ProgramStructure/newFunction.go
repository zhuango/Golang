package main

import (
	"fmt"
)

// just for syntactic conv enience
func getPointer() *int {
	return new(int)
}

// same as getPointer
func getPointerFromLocalVariable() *int {
	var a int
	return &a
}

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	p1 := new(int)
	p2 := new(int)
	fmt.Println(p1 == p2)
}