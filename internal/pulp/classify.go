package pulp

import (
	"errors"
	"fmt"
	"strings"
)

var ErrRetryable = errors.New("pulppress retryable")

func WrapRetryable(op string, err error) error {
	if err == nil {
		return nil
	}
	if strings.TrimSpace(op) == "" {
		op = "op"
	}
	return fmt.Errorf("%s: %w", op, ErrRetryable)
}

func IsRetryable(err error) bool {
	return errors.Is(err, ErrRetryable)
}

func ClassifyOutcome(err error) string {
	if err == nil {
		return "ok"
	}
	if IsRetryable(err) {
		return "retry"
	}
	return "terminal"
}
