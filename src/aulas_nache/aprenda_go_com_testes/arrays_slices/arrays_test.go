package arrays_test

import (
	"arrays"
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	t.Run("coleção de 5 números", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := arrays.Soma(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, dado %v", got, want, numbers)
		}

	})

}

func TestSomaTudo(t *testing.T) {
	want := []int{3, 9}
	got := arrays.SomaTudo([]int{1, 2}, []int{0, 9})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestSomaTodoOResto(t *testing.T) {

	t.Run("faz as somas de alguns slices", func(t *testing.T) {
		want := []int{2, 9}
		got := arrays.SomaTodoOResto([]int{1, 2}, []int{0, 9})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		want := []int{0, 9}
		got := arrays.SomaTodoOResto([]int{}, []int{3, 4, 5})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

/*func TestSomaTodoOResto(t *testing.T) {

	verificaSomas := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	}

	t.Run("faz a soma do resto", func(t *testing.T) {
		want := []int{2,9}
		got := SomaTodoOResto([]int{1,2}, []int{0,9})


		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		want := []int{0,9}
		got := SomaTodoOResto([]int{}, []int{3,4,5})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}*/
