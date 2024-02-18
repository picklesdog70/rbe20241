package domains

type Cliente struct {
	Saldo      int64
	Limite     int64
	Transacoes []Transacao
}
