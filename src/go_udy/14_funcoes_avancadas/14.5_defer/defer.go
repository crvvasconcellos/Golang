package main

import "fmt"

func f1() {
	fmt.Println("Executando a f1")
}

func f2() {
	fmt.Println("Executando a f2")
}

func AlunoAprovado(n1, n2 float64) bool {
	defer fmt.Println("Média calculada. Resultado retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	if (n1+n2)/2 > 6 {
		return true
	}
	return false
}

func main() {
	/*defer f1()
	// DEFER = adiar
	f2()*/

	fmt.Println(AlunoAprovado(7, 8))
}
