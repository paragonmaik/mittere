package errs

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidExt    = errors.New("unsupported file extension")
	ErrInvalidMethod = errors.New("unsupported HTTP method")
	ErrInvalidUrl    = errors.New("unsupported HTTP method")
)

type ReadErr struct {
	Step  string
	Msg   string
	Cause error
}

func (s *ReadErr) Error() string {
	return fmt.Sprintf("Step: %q: %s: Cause: %v", s.Step,
		s.Msg, s.Cause)
}

func (s *ReadErr) Is(target error) bool {
	t, ok := target.(*ReadErr)
	if !ok {
		return false
	}

	return t.Step == s.Step
}

func (s *ReadErr) Unwrap() error {
	return s.Cause
}
