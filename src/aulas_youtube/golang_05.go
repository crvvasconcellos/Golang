package main

import (
	"fmt"
)

func main() {

	fmt.Println(somaTudo(2, 3, 5, 20, 50))
}

func soma(a, b int) (result int) {

	result = a + b
	return
}

func somaTudo(x ...int) int {
	resultado := 0
	for _, v := range x {
		resultado += v
	}
	return resultado
}
