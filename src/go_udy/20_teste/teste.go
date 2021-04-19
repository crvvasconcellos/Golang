package main

import (
	"errors"
	"fmt"
)

var ErrExpected = errors.New("É impossivel dividir por ZERO - Resultado Indefinido")
var count int = 0

func Divide(a, b float64, c ...float64) (float64, error) {
	if b == 0 {
		return 0, ErrExpected
	}

	//count := 0
	result := a / b
	for _, i := range c {
		result /= i
		if i == 0 {
			return 0, ErrExpected
		} else {
			count += 1
		}
	}

	return result, nil
}

func main() {

	/*var value string

	fmt.Println("Enter values: ")
	fmt.Scan(&value)

	fmt.Println(value)

	for _, valor := range value {
		fmt.Println(string(valor))
	}*/

	fmt.Println(Divide(100.0, 0.1, 2.0, 1.0, 0.1))
	fmt.Println(count)

	if b != 0 && count > 0 {
		if err != nil {
			t.Fatalf("Não era esperado um erro %v, porém ele ocorreu!", err)
		}
	}

}
