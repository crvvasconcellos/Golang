package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %q no banco de dados \n", u.nome)
}

func (u *usuario) fazerNiver() {
	u.idade++
}

func main() {

	usuario1 := usuario{"Jonas", 17}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario1.fazerNiver()
	fmt.Println(usuario1.idade)
}
