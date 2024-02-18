package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	domains "github.com/picklesdog70/rbe-api/internal/core/domain"
	portsOut "github.com/picklesdog70/rbe-api/internal/core/port/out"
	"github.com/picklesdog70/rbe-api/pkg/config"
)

type transacaoRepository struct {
	dbConn *sql.DB
}

func NewTransacaoRepository(dbConn *sql.DB) portsOut.TransacaoRepository {
	return &transacaoRepository{
		dbConn: dbConn,
	}
}

func (tr *transacaoRepository) Save(ctx context.Context, transacao domains.Transacao, clienteId int64) (*domains.Cliente, error) {

	var retornoTransacaoCriada string
	criaTransacao := tr.dbConn.QueryRow("SELECT CREATE_TRANSACATION($1, $2, $3)", clienteId, transacao.Valor, transacao.Descricao)
	errCriaTransacao := criaTransacao.Scan(&retornoTransacaoCriada)

	if errCriaTransacao != nil {
		return nil, errCriaTransacao
	}

	saldo, limite, errCodigo := trataRetorno(retornoTransacaoCriada)

	fmt.Println("Estatisticas de conexão")
	fmt.Printf("Conexões abertas: %d\n", tr.dbConn.Stats().OpenConnections)
	fmt.Printf("Conexões atualmente em uso: %d\n", tr.dbConn.Stats().InUse)
	fmt.Printf("Conexões inativas: %d\n", tr.dbConn.Stats().Idle)

	switch errCodigo {
	case 0:
		return &domains.Cliente{Saldo: saldo, Limite: limite}, nil
	case -1:
		return nil, &config.NotFoundError{Id: clienteId, Entidade: "Transacoes"}
	case -2:
		return nil, &config.BusinessError{Message: "Operação deixa o saldo inconsistente"}
	case -99:
		return nil, &config.BusinessError{Message: "Erro ao converter o retorno da função"}
	default:
		return nil, &config.BusinessError{Message: "Erro geral"}
	}

}

func trataRetorno(valoresRetornado string) (saldo int64, limite int64, errCodigo int64) {
	var rgx = regexp.MustCompile(`\((.*?)\)`)
	valoresSemParenteses := rgx.FindStringSubmatch(valoresRetornado)
	valores := strings.Split(valoresSemParenteses[1], ",")

	saldoRetornado, err := strconv.ParseInt(valores[0], 10, 64)

	if err != nil {
		return 0, 0, -99
	}

	limiteRetornado, err := strconv.ParseInt(valores[1], 10, 64)

	if err != nil {
		return 0, 0, -99
	}

	errCodigoRetornado, err := strconv.ParseInt(valores[2], 10, 64)

	if err != nil {
		return 0, 0, -99
	}

	return saldoRetornado, limiteRetornado, errCodigoRetornado
}
