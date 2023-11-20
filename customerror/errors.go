package customerror

import (
	"errors"
)

var (
	ErrValidation = errors.New("file validation failed")
)
