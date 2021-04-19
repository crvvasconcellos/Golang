package main

import "fmt"

func soma(n ...int) int {
	result := 0
	for _, i := range n {
		result += i
	}
	return result
}

func main() {
	fmt.Println(soma(1, 2, 3, 4))
}
