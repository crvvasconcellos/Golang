package maps

import (
	"errors"
)

type Dicionario map[string]string

var ErroNew = errors.New("Erro")

func (d Dicionario) Busca(key string) (string, error) {
	_, ok := d[key]
	if !ok {
		return "", ErroNew
	}
	return d[key], nil
}

func (d Dicionario) Adiciona1(key, value string) {
	d[key] = value
}

func (d Dicionario) Atualiza(key, value string) {
	d[key] = value
}

func (d *Dicionario) Deleta(key string) {
	delete(*d, key)
}
