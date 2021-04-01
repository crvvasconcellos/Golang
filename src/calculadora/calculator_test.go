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
		{
			desc: "adicao 2 + 2 = 4",
			want: 4,
			a:    2,
			b:    2,
		},
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
			desc: "divisao -0.5 / -0.5 = 0.375",
			want: 0.4,
			a:    2,
			b:    5,
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
