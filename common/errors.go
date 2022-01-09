package common

import "fmt"

type ValidationError struct {
	Code    int
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("code: %d; message: %s", v.Code, v.Message)
}
