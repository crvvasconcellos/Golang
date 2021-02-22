package inteiros

import (
	"testing"
)

func TestAdicionador(t *testing.T) {
	soma := Adiciona(2, 2)
	esperado := 4

	if resultado != esperado {
		t.Errorf("resultado '%d', esperado '%d'", resultado, esperado)
	}
}
