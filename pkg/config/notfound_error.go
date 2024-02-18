package config

import "fmt"

type NotFoundError struct {
	Id       int64
	Entidade string
	Err      error
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s não encontrado(a) com o id %d", e.Entidade, e.Id)
}
