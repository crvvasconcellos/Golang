package inteiros_test

import (
	"fmt"
	"inteiros"
	"testing"
)

func TestAdicionador(t *testing.T) {
	got := inteiros.Adiciona(2, 2)
	want := 4

	if got != want {
		t.Errorf("esperado '%d', resultado '%d'", got, want)
	}
}

func ExampleAdiciona() {
	soma := inteiros.Adiciona(2, 5)
	fmt.Println(soma)
	// Output: 7
}
