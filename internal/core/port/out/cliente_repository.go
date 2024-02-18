package ports

import domains "github.com/picklesdog70/rbe-api/internal/core/domain"

type ClienteRepository interface {
	GetByIdComTransacoes(id int64, qtdeUltimasTransacoes int64) (*domains.Cliente, error)
}
