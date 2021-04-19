package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":      "Paulo",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro_nome": "Jonas",
			"ultimo_nome":   "Santos",
		},
	}
	fmt.Println(usuario2)

}
