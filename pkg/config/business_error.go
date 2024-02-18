package config

import "fmt"

type BusinessError struct {
	Message string
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf(e.Message)
}
