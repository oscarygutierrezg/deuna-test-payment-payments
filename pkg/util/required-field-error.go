package util

import "fmt"

type RequiredFieldError struct {
	Message string
}

func (e *RequiredFieldError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
