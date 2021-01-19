package main

import "fmt"

func add(s float32) (x, z float32) {
	x = s + 10
	z = x - 5
	return

}

func main() {
	fmt.Print(add(100))
}
