package main

import "fmt"

func main() {

	var ar1 [5]int
	ar1[0] = 2
	fmt.Println(ar1)

	ar2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(ar2)

	sl1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(sl1)
	fmt.Printf("%T, %T", ar2, sl1)
	fmt.Println("---------")

	sl1 = append(sl1, 57)
	fmt.Println(sl1)

	// Arrays Internos
	sl2 := make([]float64, 10, 11) // 2º valor = tamanho(lenght) // 3º valor = capacidade(cap)
	fmt.Println(sl2)
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))

	sl3 := append(sl2, 57)
	sl3 = append(sl3, 59)

	fmt.Println("---------")
	fmt.Println(sl3)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

}
