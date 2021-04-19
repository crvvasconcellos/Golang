package main

import "fmt"

func recover1() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada!")
	}
}

func AlunoAprovado(n1, n2 float64) bool {
	defer recover1()
	//fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	if (n1+n2)/2 > 6 {
		return true
	} else if (n1+n2)/2 < 6 {
		return false
	}
	panic("A média é exatamente 6")
}

func main() {
	fmt.Println(AlunoAprovado(8, 4))
}
