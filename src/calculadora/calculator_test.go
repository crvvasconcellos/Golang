package calculator_test

import (
	"calculator"
	"math/rand"
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
		{
			desc: "subtracao 0.15 - (-1) = 1.15",
			want: 1.15,
			a:    0.15,
			b:    -1,
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
		{
			desc: "multiplicacao 900 * 1500 = 1350000",
			want: 1350000,
			a:    900,
			b:    1500,
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
		{
			desc: "divisao por 0",
			want: 0,
			a:    100,
			b:    0,
		},
	}
	for _, tC := range testCases {

		t.Run(tC.desc, func(t *testing.T) {
			got, err := calculator.Divide(tC.a, tC.b)

			if tC.b == 0 {
				if err != calculator.ErrExpected {
					t.Fatalf("Esperava-se um erro %v, mas recebeu o %v", calculator.ErrExpected, err)
				}
			}

			if tC.b != 0 {
				if err != nil {
					t.Fatalf("Não era esperado um erro %v, porém ele ocorreu!", err)
				}
			}

			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
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
		{"Raiz negativa - Erro", 0, -49},
		{"Raiz quadrada 0 = 0", 0, 0},
		{"Raiz quadrada 144 = 12", 12, 144},
	}
	for _, tC := range testCases {

		t.Run(tC.desc, func(t *testing.T) {
			got, err := calculator.Sqrt(tC.a)

			if tC.a < 0 {
				if err != calculator.ErrExpected1 {
					t.Fatalf("Esperava-se um erro %v, mas recebeu o %v", calculator.ErrExpected1, err)
				}
			}

			if tC.a >= 0 {
				if err != nil {
					t.Fatalf("Não era esperado um erro %v, porém ele ocorreu!", err)
				}
			}

			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}

		})

	}

}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.NormFloat64()
		b := rand.NormFloat64()
		var want float64 = a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}
	}
}
