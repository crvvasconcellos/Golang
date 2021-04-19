package main

import "fmt"

type usuario struct {
	nome     string
	idade    int
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {
	var u usuario
	u.nome = "Davi"
	u.idade = 15

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	u2 := usuario{"Jonas", 16, enderecoExemplo}

	u3 := usuario{nome: "Maria"}

	fmt.Println(u, u2, u3)
}
