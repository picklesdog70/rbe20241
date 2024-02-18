package ports

import (
	"context"

	domains "github.com/picklesdog70/rbe-api/internal/core/domain"
)

type ClienteComTransacoesUseCase interface {
	Execute(ctx context.Context, id int64, qtdeUltimasTransacoes int64) (*domains.Cliente, error)
}
