package main

import "testing"

func TestSoma(t *testing.T) {

	t.Run("coleção de 5 números", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Soma(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, dado %v", got, want, numbers)
		}

	})

}

func TestSomaTudo(t *testing.T) {
	want := []int{3, 9}
	got := SomaTudo([]int{1, 2}, []int{0, 9})

	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}
