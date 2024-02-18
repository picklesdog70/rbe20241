package usecases

import (
	"context"

	domains "github.com/picklesdog70/rbe-api/internal/core/domain"
	portsIn "github.com/picklesdog70/rbe-api/internal/core/port/in"
	portsOut "github.com/picklesdog70/rbe-api/internal/core/port/out"
)

type clienteComTransacoesUseCase struct {
	repository portsOut.ClienteRepository
}

func NewClienteComTransacoesUseCase(repository portsOut.ClienteRepository) portsIn.ClienteComTransacoesUseCase {
	return &clienteComTransacoesUseCase{
		repository: repository,
	}
}

func (uc *clienteComTransacoesUseCase) Execute(ctx context.Context, id int64, qtdeUltimasTransacoes int64) (*domains.Cliente, error) {
	return uc.repository.GetByIdComTransacoes(id, qtdeUltimasTransacoes)
}
