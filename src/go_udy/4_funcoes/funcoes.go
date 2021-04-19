package main

import "fmt"

func Somar(a, b float64) float64 {
	return a + b
}

func CalculosMatematicos(a, b float64) (float64, float64, float64, float64) {
	soma := a + b
	subtracao := a - b
	divisao := a / b
	multiplicacao := a * b

	return soma, subtracao, divisao, multiplicacao
}

func main() {
	fmt.Println(Somar(1, 5))
	fmt.Println(CalculosMatematicos(1, 5))
}
