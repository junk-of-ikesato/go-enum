package main

import (
	"fmt"
)

const (
	A = iota
	B
	C
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	var a int = A
	fmt.Println(a)
	a = B
	fmt.Println(a)
}
