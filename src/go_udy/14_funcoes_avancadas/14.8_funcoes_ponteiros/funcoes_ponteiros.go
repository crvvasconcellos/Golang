package main

import "fmt"

func inverterSinal(n int) int {
	return n * -1
}

func inverterSinalComPonteiro(n *int) {
	*n = *n * -1
}

func main() {

	n := 20
	nInvertido := inverterSinal(n)
	fmt.Println(nInvertido)

	novo_n := 40
	fmt.Println(novo_n)
	inverterSinalComPonteiro(&novo_n)
	fmt.Println(novo_n)

	novo_n = 50
	inverterSinalComPonteiro(&novo_n)
	fmt.Println(novo_n)
}
