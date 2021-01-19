package main

import "fmt"

func add(s float32) (float32, float32) {
	x := s + 10
	z := x - 5
	return x, z

}

func main() {
	fmt.Print(add(100))
}
