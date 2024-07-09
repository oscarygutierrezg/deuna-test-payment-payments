package util

import "fmt"

type JWTError struct {
	Message string
}

func (e *JWTError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
