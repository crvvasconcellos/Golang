package main

import "fmt"

func main() {

	n := 10

	p := &n

	*p = 50

	if n > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if n2 := n; n2 > 0 {
		fmt.Println("Maior que 0")
	}

	fmt.Println(n)
}
