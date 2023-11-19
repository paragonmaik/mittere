package customerror

import (
	"errors"
	// "fmt"
)

var (
	ErrValidation = errors.New("file validation failed")
)
