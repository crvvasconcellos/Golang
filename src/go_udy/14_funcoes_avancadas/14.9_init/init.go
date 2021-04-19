package main

import "fmt"

var n int

func init() {
	fmt.Println("func init")
	n = 10
}

func main() {
	fmt.Println("func main")
	fmt.Println(n)
}
