package main

import "fmt"

func main() {

	func() {
		fmt.Println("Hello!")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Hello!")
}
