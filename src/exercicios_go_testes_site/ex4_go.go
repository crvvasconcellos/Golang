package main

import "fmt"

const quantidadeRepeticoes = 5

func Repetir(caractere string) string {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
		fmt.Println(caractere, repeticoes)
	}
	return repeticoes
}

func main() {
	fmt.Println(Repetir("c"))
}
