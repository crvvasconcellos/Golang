package ponteiros_erros

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(qtde Bitcoin) {
	//fmt.Printf("O enderço do saldo no Depositar é %v \n", &c.saldo)
	c.saldo += qtde
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

/*func (c *Carteira) Retirar(qtde Bitcoin) {
	c.saldo -= qtde
}

func (c *Carteira) Retirar(qtde Bitcoin) error {
	if qtde > c.saldo {
		return errors.New("não é possível retirar: saldo insuficiente")
	}
	c.saldo -= qtde
	return nil
}*/

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

func (c *Carteira) Retirar(qtde Bitcoin) error {

	if qtde > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= qtde
	return nil
}
