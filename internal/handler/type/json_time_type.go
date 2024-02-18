package types

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02T15:04:05.000000Z"))), nil
}
