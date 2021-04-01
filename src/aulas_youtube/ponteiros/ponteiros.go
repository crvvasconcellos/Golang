package main

import "fmt"

func main() {

	//memória---0xc0000120a0(50), <----- a <----- 50

	a := 50

	ponteiros := &a

	*ponteiros = 100

	b := &a

	*b = 200

	fmt.Println(a)
}
