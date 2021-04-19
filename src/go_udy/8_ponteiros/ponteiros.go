package main

import "fmt"

func main() {

	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)

	v1++
	fmt.Println(v1, v2)

	// Ponteiro é uma referência de memória
	var v3 int = 100
	var ponteiro *int
	ponteiro = &v3
	fmt.Println(v3, ponteiro)
	fmt.Println(v3, *ponteiro) // desreferenciamento

}
