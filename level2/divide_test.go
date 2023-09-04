package level2_test

import (
	"errors"
	"testing"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide: cannot divide by 0")
	}
	return a / b, nil
}

func TestDivideByZero(t *testing.T) {
	_, err := divide(1, 0)
	if err == nil {
		t.Fatal("err was nil")
	}
	t.Log(err)
}
