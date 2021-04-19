package main

import "fmt"

func DiaDaSemana(n int) string {
	switch n {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda_Feira"
	default:
		return "numero invalido"
	}

}

func DiaDaSemana2(n int) string {
	switch {
	case n == 1:
		return "Domingo"
	default:
		return "Erroooouuu"
	}
}

func DiaDaSemana3(n int) string {
	var dia string
	switch {
	case n == 1:
		dia = "Domingo"
	default:
		dia = "Erroooouuu"
	}
	return dia
}

func main() {
	fmt.Println(DiaDaSemana(20))

	fmt.Println(DiaDaSemana2(20))

	fmt.Println(DiaDaSemana3(20))

}
