package level4_test

import (
	"errors"
	"testing"
)

type MultiError []error

func (e MultiError) Error() string {
	msg := ""
	for _, err := range e {
		msg += err.Error() + ";"
	}
	return msg
}

func TestMultiError(t *testing.T) {
	var err error = MultiError{
		errors.New("err1"),
		errors.New("err2"),
		errors.New("err3"),
	}
	var multiErrs MultiError
	if !errors.As(err, &multiErrs) {
		t.Fatalf("should have been MultiError")
	}
	if len(multiErrs) != 3 {
		t.Fatalf("expected 3 errors")
	}
}
