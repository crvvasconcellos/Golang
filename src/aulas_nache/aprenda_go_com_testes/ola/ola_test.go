package main

import "testing"

func TestOla(t *testing.T) {
	verificarMensagemCorreta := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("em frances", func(t *testing.T) {
		got := Ola("teste ", "frances")
		want := "Bonjour, teste"
		verificarMensagemCorreta(t, got, want)

	})

}

//t.Run("diz 'olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
//got := Ola("")
//want := "Olá, mundo"

//if got != want {
//t.Errorf("got %q, want %q", got, want)
//}
//})
