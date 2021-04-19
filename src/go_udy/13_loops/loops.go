package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		time.Sleep(time.Millisecond)
		fmt.Println(i)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Millisecond)
	}

	nomes := [3]string{"Jonas", "Caio", "Carla"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for ind, letra := range "palavra" {
		fmt.Println(ind, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Josias",
		"sobrenome": "boneco",
	}

	for chave, name := range usuario {
		fmt.Println(chave, name)
	}

}
