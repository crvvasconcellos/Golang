package main

import "fmt"

type pessoa struct {
	nome   string
	idade  int
	altura float64
}

type estudante struct {
	pessoa
	curso string
	reg   int
}

func main() {
	p1 := pessoa{"Jonas", 20, 1.78}

	fmt.Println(p1)

	est1 := estudante{p1, "engenharia", 123454}

	fmt.Println(est1.altura)
}
