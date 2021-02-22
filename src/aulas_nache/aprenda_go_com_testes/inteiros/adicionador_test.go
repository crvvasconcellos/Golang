package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	got := Adiciona(2, 2)
	want := 4

	if got != want {
		t.Errorf("esperado '%d', resultado '%d'", got, want)
	}
}

func ExampleAdiciona() {
	soma := Adiciona(2, 5)
	fmt.Println(soma)
	// Output: 7
}
