package main

import "fmt"

func teste(s int) (int, int) {
	x := s * 4 / 5
	y := s - x
	return x, y
}

func main() {
	fmt.Println(teste(26))
}
