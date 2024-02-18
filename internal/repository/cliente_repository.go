package repositories

import (
	"database/sql"

	domains "github.com/picklesdog70/rbe-api/internal/core/domain"
	ports "github.com/picklesdog70/rbe-api/internal/core/port/out"
	"github.com/picklesdog70/rbe-api/pkg/config"

	"github.com/picklesdog70/rbe-api/internal/repository/entity"
)

type clienteRepository struct {
	dbConn *sql.DB
}

func NewClienteRepository(dbConn *sql.DB) ports.ClienteRepository {
	return &clienteRepository{
		dbConn: dbConn,
	}
}

func (cr *clienteRepository) GetByIdComTransacoes(id int64, qtdeUltimasTransacoes int64) (*domains.Cliente, error) {

	var cliente entity.ClienteEntity
	clienteQuery := cr.dbConn.QueryRow("SELECT saldo, limite FROM CLIENTES WHERE id = $1", id)
	errQueryClient := clienteQuery.Scan(&cliente.Saldo, &cliente.Limite)

	if errQueryClient != nil {

		if errQueryClient == sql.ErrNoRows {
			return nil, &config.NotFoundError{Id: id, Entidade: "Clientes"}
		}

		return nil, errQueryClient
	}

	transacoesQuery, err := cr.dbConn.Query("SELECT valor, descricao, data from transacoes WHERE cliente_id = $1 ORDER BY data DESC LIMIT $2", id, qtdeUltimasTransacoes)

	if err != nil {
		return nil, err
	}

	var transacoes []entity.TransacaoEntity
	for transacoesQuery.Next() {
		var transacao entity.TransacaoEntity
		if err := transacoesQuery.Scan(
			&transacao.Valor,
			&transacao.Descricao,
			&transacao.Data); err != nil {
			return nil, err
		}

		transacoes = append(transacoes, transacao)
	}

	return toDomainComTransacoes(cliente, transacoes), nil
}

func toDomainComTransacoes(cliente entity.ClienteEntity, transacoes []entity.TransacaoEntity) *domains.Cliente {
	var ultimasTransacoesResponse []domains.Transacao

	for _, transacao := range transacoes {
		ultimasTransacoesResponse = append(ultimasTransacoesResponse, domains.Transacao{
			Valor:     transacao.Valor,
			Descricao: transacao.Descricao,
			Data:      transacao.Data,
		})
	}

	return &domains.Cliente{
		Saldo:      cliente.Saldo,
		Limite:     cliente.Limite,
		Transacoes: ultimasTransacoesResponse,
	}
}
