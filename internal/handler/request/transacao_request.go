package requests

type TransacaoRequest struct {
	Valor     int64  `json:"valor" binding:"required,gt=0"`
	Tipo      string `json:"tipo" binding:"required,oneof=c d"`
	Descricao string `json:"descricao" binding:"required,lte=10"`
}
