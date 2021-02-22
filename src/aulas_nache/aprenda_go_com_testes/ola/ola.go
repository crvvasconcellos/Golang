package main

import "fmt"

const idioma = "espanhol"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	if idioma == "espanhol" {
		return prefixoOlaEspanhol + nome
	}

	if idioma == "frances" {
		return prefixoOlaFrances + nome
	}

	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("", "frances"))
}
