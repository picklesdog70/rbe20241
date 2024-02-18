package ports

import (
	"context"

	domains "github.com/picklesdog70/rbe-api/internal/core/domain"
)

type TransacaoRepository interface {
	Save(ctx context.Context, transacao domains.Transacao, clienteId int64) (*domains.Cliente, error)
}
