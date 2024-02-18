package responses

import types "github.com/picklesdog70/rbe-api/internal/handler/type"

type ExtratoResponse struct {
	Saldo             ExtratoSaldoResponse       `json:"saldo"`
	UltimasTransacoes []ExtratoTransacaoResponse `json:"ultimas_transacoes,omitempty"`
}

type ExtratoSaldoResponse struct {
	Total       int64          `json:"total"`
	DataExtrato types.JSONTime `json:"data_extrato"`
	Limite      int64          `json:"limite"`
}

type ExtratoTransacaoResponse struct {
	Valor     int64          `json:"valor"`
	Tipo      string         `json:"tipo"`
	Descricao string         `json:"descricao"`
	Data      types.JSONTime `json:"realizada_em"`
}
