package level2_test

import (
	"fmt"
	"testing"
)

func subtract(a int, b int) (int, error) {
	result := a - b
	if result < 0 {
		return result, fmt.Errorf("level2: %d - %d is a negative number", a, b)
	}
	return result, nil
}

func TestSubtractNegativeError(t *testing.T) {
	_, err := subtract(69, 420)
	if err == nil {
		t.Fatal("expected subtract to return error")
	}
	t.Log(err)
}
