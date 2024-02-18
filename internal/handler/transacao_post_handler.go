package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	domains "github.com/picklesdog70/rbe-api/internal/core/domain"
	ports "github.com/picklesdog70/rbe-api/internal/core/port/in"
	requests "github.com/picklesdog70/rbe-api/internal/handler/request"
	responses "github.com/picklesdog70/rbe-api/internal/handler/response"
	"github.com/picklesdog70/rbe-api/internal/handler/util"
	"github.com/picklesdog70/rbe-api/pkg/config"
)

type TransacaoPostHandler struct {
	transacaoSaveUseCase ports.TransacaoSaveUseCase
}

func NewTransacaoPostHandler(transacaoSaveUseCase ports.TransacaoSaveUseCase) *TransacaoPostHandler {
	return &TransacaoPostHandler{
		transacaoSaveUseCase: transacaoSaveUseCase,
	}
}

func (hdl *TransacaoPostHandler) PostTransacao(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	var transacaoRequest requests.TransacaoRequest

	if err := util.BindJSON(&transacaoRequest, c); err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	clienteAtualizado, err := hdl.transacaoSaveUseCase.Execute(c, toDomain(transacaoRequest), id)

	if err != nil {

		var notfoundError *config.NotFoundError
		if errors.As(err, &notfoundError) {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		var businessError *config.BusinessError
		if errors.As(err, &businessError) {
			c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, toResponse(clienteAtualizado))

}

func toResponse(transacaoSalva *domains.Cliente) responses.ClienteResponse {
	return responses.ClienteResponse{
		Limite: transacaoSalva.Limite,
		Saldo:  transacaoSalva.Saldo,
	}
}

func toDomain(transacaoRequest requests.TransacaoRequest) domains.Transacao {

	var valor int64
	if transacaoRequest.Tipo == "d" {
		valor = transacaoRequest.Valor * -1
	} else {
		valor = transacaoRequest.Valor
	}

	return domains.Transacao{
		Valor:     valor,
		Descricao: transacaoRequest.Descricao,
	}
}
