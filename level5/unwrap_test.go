package level5_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnwrap(t *testing.T) {
	err := fmt.Errorf("error encountered: %w", errors.New("something bad"))
	t.Log(err)
	underlyingErr := errors.Unwrap(err)
	t.Log(underlyingErr)
}
