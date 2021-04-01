package ponteiros_erros

import (
	"fmt"
	"testing"
)

/*func TestCarteira(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(10)

	resultado := carteira.Saldo()
	esperado := 10

	if resultado != esperado {
		t.Errorf("resultado %d, esperado %d", resultado, esperado)
	}
}

func TestCarteira2(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(10)
	got := carteira.Saldo()
	//fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)
	want := 10

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}*/

func TestCarteira3(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(Bitcoin(10))
	got := carteira.Saldo()
	fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)
	want := Bitcoin(10)

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCarteira4(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(Bitcoin(10))
	got := carteira.Saldo()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestCarteira5(t *testing.T) {

	t.Run("depositar", func(t *testing.T) {
		carteira := Carteira{}

		carteira.Depositar(Bitcoin(10))
		got := carteira.Saldo()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}

		carteira.Retirar(Bitcoin(10))
		got := carteira.Saldo()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

	})

}

func TestCarteira6(t *testing.T) {

	confirmaSaldo := func(t *testing.T, carteira Carteira, want Bitcoin) {
		t.Helper()
		got := carteira.Saldo()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))

	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		carteira.Retirar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))

	})

}

func TestCarteira7(t *testing.T) {

	confirmaSaldo := func(t *testing.T, carteira Carteira, want Bitcoin) {
		t.Helper()
		got := carteira.Saldo()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))

	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		carteira.Retirar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))

	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)

		if erro == nil {
			t.Error("esperava um erro, mas nenhum ocorreu")
		}

	})

}

func TestCarteira8(t *testing.T) {

	confirmaSaldo := func(t *testing.T, carteira Carteira, want Bitcoin) {
		t.Helper()
		got := carteira.Saldo()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))

	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		carteira.Retirar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))

	})

	confirmaErro := func(t *testing.T, valor error, want string) {
		t.Helper()
		if valor == nil {
			t.Fatal("esperava um erro, mas nenhum ocorreu")
		}

		got := valor.Error()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaErro(t, erro, "não é possível retirar: saldo insuficiente")

	})

}

func TestCarteira9(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))

		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Retirar com saldo suficiente", func(t *testing.T) {
		carteira := Carteira{Bitcoin(20)}
		erro := carteira.Retirar(Bitcoin(10))

		confirmaSaldo(t, carteira, Bitcoin(10))
		confirmaErroInexistente(t, erro)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente)
	})
}

func confirmaSaldo(t *testing.T, carteira Carteira, want Bitcoin) {
	t.Helper()
	got := carteira.Saldo()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func confirmaErroInexistente(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("erro inesperado recebido")
	}
}

func confirmaErro(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("esperava um erro, mas nenhum ocorreu")
	}

	if got != want {
		t.Errorf("erro got %s, erro want %s", got, want)
	}
}
