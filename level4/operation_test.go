package level4_test

import (
	"errors"
	"fmt"
	"testing"
)

type UnsupportedModeError struct {
	mode int
}

func (e *UnsupportedModeError) Error() string {
	return fmt.Sprintf("operation: mode %d is unsupported", e.mode)
}

var ErrUnsupported = errors.New("this mode is unsupported right now")

func operation(mode int) (string, error) {
	if mode < 0 {
		return "", &UnsupportedModeError{mode: mode}
	}
	return fmt.Sprintf("running operation with mode %d", mode), nil
}

func TestOperationUnsupported(t *testing.T) {
	_, err := operation(-1)
	if err == nil {
		t.Fatal("expected error for unsupported operation")
	}
	var unsupportedErr *UnsupportedModeError
	if errors.As(err, &unsupportedErr) {
		if unsupportedErr.mode != -1 {
			t.Fatalf("oh no!")
		}
	}
}
