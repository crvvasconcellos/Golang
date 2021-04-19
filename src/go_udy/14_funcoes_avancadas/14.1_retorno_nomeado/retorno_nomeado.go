package main

import "fmt"

func CalcMatematico(a, b float64) (soma float64, subtracao float64) {
	soma = a + b
	subtracao = a - b
	return
}

func main() {
	fmt.Println(CalcMatematico(2, 4))
}
