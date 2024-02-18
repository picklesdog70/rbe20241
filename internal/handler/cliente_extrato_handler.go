package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	domains "github.com/picklesdog70/rbe-api/internal/core/domain"
	ports "github.com/picklesdog70/rbe-api/internal/core/port/in"
	responses "github.com/picklesdog70/rbe-api/internal/handler/response"
	types "github.com/picklesdog70/rbe-api/internal/handler/type"
	"github.com/picklesdog70/rbe-api/pkg/config"
)

type ClienteExtratoHandler struct {
	clienteComTransacoesUseCase ports.ClienteComTransacoesUseCase
}

func NewClienteExtratoHandler(clienteComTransacoesUseCase ports.ClienteComTransacoesUseCase) *ClienteExtratoHandler {
	return &ClienteExtratoHandler{
		clienteComTransacoesUseCase: clienteComTransacoesUseCase,
	}
}

func (hld *ClienteExtratoHandler) GetExtratoByClientId(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	clienteComUltimasTransacoes, err := hld.clienteComTransacoesUseCase.Execute(c, id, 10)

	if err != nil {

		var notfoundError *config.NotFoundError
		if errors.As(err, &notfoundError) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, toExtratoResponse(clienteComUltimasTransacoes))

}

func toExtratoResponse(cliente *domains.Cliente) responses.ExtratoResponse {

	var ultimasTransacoes []responses.ExtratoTransacaoResponse

	for _, transacao := range cliente.Transacoes {

		var valor int64
		var tipo string
		if transacao.Valor < 0 {
			valor = transacao.Valor * -1
			tipo = "d"
		} else {
			valor = transacao.Valor
			tipo = "c"
		}

		ultimasTransacoes = append(ultimasTransacoes, responses.ExtratoTransacaoResponse{
			Valor:     valor,
			Tipo:      tipo,
			Descricao: transacao.Descricao,
			Data:      types.JSONTime(transacao.Data),
		})
	}

	return responses.ExtratoResponse{
		Saldo: responses.ExtratoSaldoResponse{
			Total:       cliente.Saldo,
			Limite:      cliente.Limite,
			DataExtrato: types.JSONTime(time.Now()),
		},
		UltimasTransacoes: ultimasTransacoes,
	}
}
