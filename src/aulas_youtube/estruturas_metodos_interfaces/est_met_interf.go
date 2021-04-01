package main

import "fmt"

//Struct:
type Pessoa struct {
	Nome    string
	Idade   int
	Veiculo Veiculo
}

type Carro struct {
	Fabricante string
	Modelo     string
	Ano        int
}

type Moto struct {
	Fabricante string
	Ano        int
}

// Método:
func (p Pessoa) Andou() {
	fmt.Println(p.Nome, "Andou a mais de 200 km/h no", p.Veiculo)
}

//Interface:
type Veiculo interface {
	buzina()
}

func (c Carro) buzina() {
	fmt.Println("Andou a mais de 200 km/h no", c.Modelo)
}

func main() {

	carro := Carro{
		Fabricante: "Honda",
		Modelo:     "Civic 1.7",
		Ano:        2002,
	}

	/*moto := Moto{
		Fabricante: "Ducati",
		Ano:        2015,
	}*/

	cyro := Pessoa{
		Nome:    "Cyro",
		Idade:   32,
		Veiculo: carro,
	}

	fmt.Println("Nome: ", cyro.Nome)
	fmt.Println("Idade: ", cyro.Idade)
	//fmt.Println("Carro: ", cyro.Carro.Modelo)
	fmt.Println("Carro: ", cyro.Veiculo)
	cyro.Andou()
	carro.buzina()

}
