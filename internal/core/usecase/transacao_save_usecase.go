package usecases

import (
	"context"

	domains "github.com/picklesdog70/rbe-api/internal/core/domain"
	portsIn "github.com/picklesdog70/rbe-api/internal/core/port/in"
	portsOut "github.com/picklesdog70/rbe-api/internal/core/port/out"
)

type transacaoSaveUseCase struct {
	repository portsOut.TransacaoRepository
}

func NewTransacaoSaveUseCase(repository portsOut.TransacaoRepository) portsIn.TransacaoSaveUseCase {
	return &transacaoSaveUseCase{
		repository: repository,
	}
}

func (uc *transacaoSaveUseCase) Execute(ctx context.Context, transacao domains.Transacao, clienteId int64) (*domains.Cliente, error) {
	return uc.repository.Save(ctx, transacao, clienteId)
}
