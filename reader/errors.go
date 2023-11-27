package reader

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidExt = errors.New("unsupported file extension")
)

type readErr struct {
	step  string
	msg   string
	cause error
}

func (s *readErr) Error() string {
	return fmt.Sprintf("Step: %q: %s: Cause: %v", s.step,
		s.msg, s.cause)
}

func (s *readErr) Is(target error) bool {
	t, ok := target.(*readErr)
	if !ok {
		return false
	}

	return t.step == s.step
}

func (s *readErr) Unwrap() error {
	return s.cause
}
