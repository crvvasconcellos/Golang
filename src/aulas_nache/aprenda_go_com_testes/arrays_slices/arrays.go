package main

func Soma(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SomaTudo(numbers, numbers2 []int) []int {
	sum := 0
	sum2 := 0
	for _, number := range numbers {
		sum += number
	}
	for _, number2 := range numbers2 {
		sum2 += number2
	}
	return []int{sum, sum2}
}
