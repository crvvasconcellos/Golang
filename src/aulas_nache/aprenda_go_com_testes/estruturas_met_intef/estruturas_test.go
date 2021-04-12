package estruturas_test

import (
	"estruturas"
	"testing"
)

func TestPerimetro(t *testing.T) {
	retangulo := estruturas.Retangulo{10.0, 10.0}
	want := 40.0
	got := estruturas.Perimetro(retangulo)

	if want != got {
		t.Errorf("want %.2f got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	retangulo := estruturas.Retangulo{12.0, 6.0}
	want := 72.0
	got := retangulo.Area()

	if want != got {
		t.Errorf("want %.2f got %.2f", want, got)
	}
}

func TestCircle(t *testing.T) {
	circle := estruturas.Circle{10.0}
	want := 314.159265
	got := estruturas.AreaCircle(circle)

	if want != got {
		t.Errorf("want %.2f got %.2f", want, got)
	}
}

func TestAreaNovo(t *testing.T) {

	t.Run("retângulos", func(t *testing.T) {
		retangulo := estruturas.Retangulo{12.0, 6.0}
		want := 72.0
		got := retangulo.Area()

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

	t.Run("círculos", func(t *testing.T) {
		circulo := estruturas.Circle{10.0}
		want := 314.159265
		got := circulo.Area()

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})
}

func TestAreaNovo2(t *testing.T) {

	verificaArea := func(t *testing.T, forma estruturas.Forma, want float64) {
		t.Helper()
		got := forma.Area()

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	}

	t.Run("retângulos", func(t *testing.T) {

		retangulo := estruturas.Retangulo{12.0, 6.0}
		verificaArea(t, retangulo, 72.0)

	})

	t.Run("círculos", func(t *testing.T) {

		circulo := estruturas.Circle{10}
		verificaArea(t, circulo, 314.159265)

	})
}

func TestAreaNovo3(t *testing.T) {
	testesArea := []struct {
		forma estruturas.Forma
		want  float64
	}{
		{estruturas.Retangulo{12, 6}, 72.0},
		{estruturas.Circle{10}, 314.159265},
	}

	for _, tt := range testesArea {
		got := tt.forma.Area()
		if got != tt.want {
			t.Errorf("want %.2f, got %.2f", tt.want, got)
		}
	}
}

func TestAreaNovo4(t *testing.T) {
	testesArea := []struct {
		forma estruturas.Forma
		want  float64
	}{
		{estruturas.Retangulo{12, 6}, 72.0},
		{estruturas.Circle{10}, 314.159265},
		{estruturas.Triangulo{12, 6}, 36.0},
	}

	for _, tt := range testesArea {
		got := tt.forma.Area()
		if got != tt.want {
			t.Errorf("want %.2f, got %.2f", tt.want, got)
		}
	}
}

func TestAreaNovo5(t *testing.T) {
	testesArea := []struct {
		forma estruturas.Forma
		want  float64
	}{
		{forma: estruturas.Retangulo{Largura: 12, Altura: 6}, want: 72.0},
		{forma: estruturas.Circle{Raio: 10}, want: 314.159265},
		{forma: estruturas.Triangulo{Base: 12, Altura: 6}, want: 36.0},
	}

	for _, tt := range testesArea {
		got := tt.forma.Area()
		if got != tt.want {
			t.Errorf("want %.2f, got %.2f", tt.want, got)
		}
	}
}

func TestAreaNovo6(t *testing.T) {
	testesArea := []struct {
		nome    string
		forma   estruturas.Forma
		temArea float64
	}{
		{nome: "Retângulo", forma: estruturas.Retangulo{Largura: 12, Altura: 6}, temArea: 72.0},
		{nome: "Círculo", forma: estruturas.Circle{Raio: 10}, temArea: 314.159265},
		{nome: "Triângulo", forma: estruturas.Triangulo{Base: 12, Altura: 6}, temArea: 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			got := tt.forma.Area()
			if got != tt.temArea {
				t.Errorf("%#v got %.2f, want %.2f", tt.forma, got, tt.temArea)
			}
		})
	}
}
