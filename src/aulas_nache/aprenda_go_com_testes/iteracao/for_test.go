package iteracao_test

import (
	"iteracao"
	"testing"
)

func TestRepetir(t *testing.T) {
	got := iteracao.Repetir("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("got '%s' but want '%s'", got, want)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteracao.Repetir("a")
	}
}
