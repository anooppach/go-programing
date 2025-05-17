package fibonacci

import (
	"testing"
)

func Test1(t *testing.T) {

	actual := Fibonacci(10)
	expected := 55
	if actual != expected {
		t.Errorf("Fibonacci(10) = %d; expected %d", actual, expected)
	}
}

func Test2(t *testing.T) {

	actual := Fibonacci(1)
	expected := 1
	if actual != expected {
		t.Errorf("Fibonacci(10) = %d; expected %d", actual, expected)
	}
}
