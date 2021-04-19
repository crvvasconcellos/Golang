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

func TestAdd(t *testing.T) {
	testCases := []struct {
		desc       string
		want, a, b float64
		c          []float64
	}{
		{desc: "adicao 2 + 2 = 4", want: 4, a: 2, b: 2},
		{desc: "adicao 1 + 1.25 = 2.25", want: 2.25, a: 1, b: 1.25},
		{"adicao 1 + 1.25 = 2.25", 10.25, 1, 1.25, []float64{5.0, 3.0}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Add(tC.a, tC.b, tC.c...)
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
		c          []float64
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
		{
			desc: "subtracao 0.15 - (-1) - (-100) = 101.15",
			want: 101.15,
			a:    0.15,
			b:    -1,
			c:    []float64{-100.0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Subtract(tC.a, tC.b, tC.c...)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}

		})
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		desc string
		want float64
		a    float64
		b    float64
		c    []float64
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
		{
			desc: "multiplicacao 1 * 2 * 3 = 6",
			want: 6,
			a:    1.0,
			b:    2.0,
			c:    []float64{3.0},
		},
		{
			desc: "multiplicacao 3 * 3 * 3 * 3 = 81",
			want: 81,
			a:    3.0,
			b:    3.0,
			c:    []float64{3.0, 3.0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Multiply(tC.a, tC.b, tC.c...)
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
		c          []float64
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
		{
			desc: "divisao 100 / 10 / 5 / 2 = 1",
			want: 1,
			a:    100,
			b:    10,
			c:    []float64{5.0, 2.0},
		},
		{
			desc: "divisao 100 / 10 / 5 / 2 / 1 / 0 = erro",
			want: 0,
			a:    100,
			b:    10,
			c:    []float64{5.0, 2.0, 1.0, 0.0},
		},
		{
			desc: "divisao 100 / 10 / 0 = erro",
			want: 0,
			a:    100,
			b:    10,
			c:    []float64{0.0},
		},
	}
	for _, tC := range testCases {

		t.Run(tC.desc, func(t *testing.T) {
			got, err := calculator.Divide(tC.a, tC.b, tC.c...)

			c := 0
			for _, i := range tC.c {
				if i == 0 {
					if err != calculator.ErrExpected {
						t.Fatalf("Esperava-se um erro %v, mas recebeu o %v", calculator.ErrExpected, err)
					}
				} else {
					c += 1
				}
			}

			if tC.b != 0 && c == len(tC.c) {
				if err != nil {
					t.Fatalf("Não era esperado um erro %v, porém ele ocorreu!", err)
				}
			}

			if tC.b == 0 {
				if err != calculator.ErrExpected {
					t.Fatalf("Esperava-se um erro %v, mas recebeu o %v", calculator.ErrExpected, err)
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
		{desc: "raiz quadrada 7 = 49", want: 7, a: 49},
		{"Raiz negativa - Erro", 1.4142135623730951, 2},
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
		c := rand.NormFloat64()
		var want float64 = a + b + c
		got := calculator.Add(a, b, c)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}
	}
}
