package main

import "fmt"

func main() {

	i := 1
	for i < 1000 {
		i += i
		fmt.Println(i)
	}
}
