package main

import "fmt"

func main() {
	// Aritimeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// Atribuicao

	var variavel1 string = "string"
	variavel2 := "string2"
	fmt.Println(variavel1, variavel2)

	// Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 <= 2)
	fmt.Println(2 == 2)
	fmt.Println(2 != 2)

	// Operadores Logicos
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Operadores Unários
	n := 10
	n++    // n = n + 1
	n += 1 // n = n + 1
	n--    // n = n - 1
	n -= 2 // n = n - 2
	n *= 3 // n = n * 3
	fmt.Println(n)

}
