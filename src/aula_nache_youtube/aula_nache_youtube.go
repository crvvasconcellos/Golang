package main

import (
	"fmt"
	"math/rand"
	"time"
)

const cols, rows, totalCardNumber, totalBalls = 5, 5, 24, 75

var card [totalCardNumber]int

func main() {
	fmt.Println("Bem vindo, jogador")
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= totalCardNumber; i++ {
		n := rand.Intn(totalBalls-1) + 1
		card[i-1] = n
	}
	fmt.Println(card)
}
