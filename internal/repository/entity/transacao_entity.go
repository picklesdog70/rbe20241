package entity

import "time"

type TransacaoEntity struct {
	Valor     int64
	Descricao string
	Data      time.Time
}
