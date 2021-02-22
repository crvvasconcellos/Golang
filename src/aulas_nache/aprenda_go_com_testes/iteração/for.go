package iteracao

const quantidadeRepetcoes = 5

func Repetir(caractere string) string {
	var repeticoes string
	for i := 0; i < quantidadeRepetcoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}

/* func Repetir(caractere string) string {
	return ""
}
*/
