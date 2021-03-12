package arrays

func Soma(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SomaTudo(numeroParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numeroParaSomar {
		somas = append(somas, Soma(numeros))
	}

	return somas
}

func SomaTodoOResto(numeroParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numeroParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
			continue
		}
		somas = append(somas, Soma(numeros[1:]))

	}
	return somas
}
