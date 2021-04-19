package main

import "fmt"

func main() {
	var variavel1 string = "varriavel1"

	variavel2 := "variavel2"

	var (
		variavel3 string = "3"
		variavel4 string = "4"
	)

	variavel5, variavel6 := "5", "6"

	fmt.Printf("%v, %v", variavel1, variavel2)

	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)
}
