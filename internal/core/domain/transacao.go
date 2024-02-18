package domains

import "time"

type Transacao struct {
	Valor     int64
	Descricao string
	Data      time.Time
}
