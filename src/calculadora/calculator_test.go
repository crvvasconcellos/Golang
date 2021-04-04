package calculator_test

import (
	"calculator"
	"testing"
)

func TestSubtract1(t *testing.T) {
	t.Parallel()
	var want float64 = -4
	got := calculator.Subtract(4, 8)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestMultiply1(t *testing.T) {
	t.Parallel()
	var want float64 = 42
	got := calculator.Multiply(7, 6)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		desc       string
		want, a, b float64
	}{
		{"adicao 2 + 2 = 4", 4, 2, 2},
		{"adicao 1 + 1.25 = 2.25", 2.25, 1, 1.25},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Add(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}

		})
	}
}

func TestSubstract(t *testing.T) {
	testCases := []struct {
		desc       string
		want, a, b float64
	}{
		{
			desc: "subtracao 2 - 16 = -14",
			want: -14,
			a:    2,
			b:    16,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Subtract(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}

		})
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		desc       string
		want, a, b float64
	}{
		{
			desc: "multiplicacao -0.5 * -0.75 = 0.375",
			want: 0.375,
			a:    -0.5,
			b:    -0.75,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Multiply(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}

		})
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		desc       string
		want, a, b float64
	}{
		{
			desc: "divisao -0.5 / -2 = 0.25",
			want: 0.25,
			a:    -0.5,
			b:    -2,
		},
	}
	for _, tC := range testCases {

		t.Run(tC.desc, func(t *testing.T) {
			got, _ := calculator.Divide(tC.a, tC.b)

			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}

		})

		t.Run("Divisão por 0", func(t *testing.T) {
			_, err := calculator.Divide(tC.a, tC.b)

			if err != nil && err == calculator.ErrExpected {
				t.Fatal(calculator.ErrExpected)
			}

		})

	}

}

func TestSqrt(t *testing.T) {
	testCases := []struct {
		desc    string
		want, a float64
	}{
		{"raiz quadrada 7 = 49", 7, 49},
		//{"Raiz negativa - Erro", 0, -49},
		{"Raiz quadrada 0 = 0", 0, 0},
		{"Raiz quadrada 144 = 12", 12, 144},
	}
	for _, tC := range testCases {

		t.Run(tC.desc, func(t *testing.T) {
			got, _ := calculator.Sqrt(tC.a)

			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}

		})

		t.Run("Raíz negativa", func(t *testing.T) {
			_, err := calculator.Sqrt(tC.a)

			if err != nil && err == calculator.ErrExpected1 {
				t.Fatal(calculator.ErrExpected1)
			}

		})

	}

}
