// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

var ErrExpected = errors.New("É impossivel dividir por ZERO - Resultado Indefinido")
var ErrExpected1 = errors.New("Resultado Indefinido")

/*type Calc struct {
	a, b float64
}*/

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64, c ...float64) float64 {
	result := a + b
	for _, i := range c {
		result += i
	}
	return result
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64, c ...float64) float64 {
	result := a - b
	for _, i := range c {
		result -= i
	}
	return result
}

func Multiply(a, b float64, c ...float64) float64 {
	result := a * b
	for _, i := range c {
		result *= i
	}
	return result
}

func Divide(a, b float64, c ...float64) (float64, error) {
	if b == 0 {
		return 0, ErrExpected
	}

	result := a / b
	for _, i := range c {
		result /= i
		if i == 0 {
			return 0, ErrExpected
		}
	}

	return result, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrExpected1
	}
	return sqrt(a), nil
}

func sqrt(a float64) float64 {
	//fmt.Println(math.Pow(a, 0.5))
	return math.Pow(a, 0.5)
}

/*func main() {
	fmt.Println(Sqrt(0))
}*/
