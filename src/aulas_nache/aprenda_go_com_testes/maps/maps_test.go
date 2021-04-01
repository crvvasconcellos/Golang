package maps

import (
	"testing"
)

func TestBusca2(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		got, _ := dicionario.Busca("teste")
		want := "isso é apenas um teste"

		if got != want {
			t.Errorf("got '%s', want '%s', dado '%s'", got, want, "test")
		}
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, err := dicionario.Busca("desconhecida")

		if err == nil {
			t.Fatal("é esperado que um erro seja obtido.")
		}
	})
}

func TestAdiciona(t *testing.T) {
	dicionario := Dicionario{}
	dicionario.Adiciona1("teste", "isso é apenas um teste")

	got := "isso é apenas um teste"
	want, err := dicionario.Busca("teste")
	if err != nil {
		t.Fatal("não foi possível encontrar a palavra adicionada:", err)
	}

	if got != want {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestUpdate(t *testing.T) {
	palavra := "teste"
	definicao := "isso é apenas um teste"
	dicionario := Dicionario{palavra: definicao}
	novaDefinicao := "nova definição"

	dicionario.Atualiza(palavra, novaDefinicao)

}

func TestDeleta(t *testing.T) {
	palavra := "teste"
	dicionario := Dicionario{palavra: "definição de teste"}

	dicionario.Deleta(palavra)

	_, err := dicionario.Busca(palavra)
	if err != ErroNew {
		t.Errorf("espera-se que '%s' seja deletado", palavra)
	}
}
