package pkg

import (
	usecases "github.com/picklesdog70/rbe-api/internal/core/usecase"
	handlers "github.com/picklesdog70/rbe-api/internal/handler"
	repositories "github.com/picklesdog70/rbe-api/internal/repository"
	"github.com/picklesdog70/rbe-api/pkg/config"
)

type Dependency struct {
	ClienteExtratoHandler *handlers.ClienteExtratoHandler
	TransacaoPostHandler  *handlers.TransacaoPostHandler
}

func NewDependency() (*Dependency, error) {

	dbConn, err := config.DbConnection()

	if err != nil {
		return nil, err
	}

	transactionRepository := repositories.NewTransacaoRepository(dbConn)
	clienteRepository := repositories.NewClienteRepository(dbConn)

	clienteComTransacoesUseCase := usecases.NewClienteComTransacoesUseCase(clienteRepository)
	transacaoSaveUseCase := usecases.NewTransacaoSaveUseCase(transactionRepository)

	return &Dependency{
		ClienteExtratoHandler: handlers.NewClienteExtratoHandler(clienteComTransacoesUseCase),
		TransacaoPostHandler:  handlers.NewTransacaoPostHandler(transacaoSaveUseCase),
	}, nil
}
