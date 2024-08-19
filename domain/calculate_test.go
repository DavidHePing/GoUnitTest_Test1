package domain

import "testing"

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	expect := 3

	if Add(a, b) != expect {
		t.Errorf("Expected answer is %v", expect)
	}
}

func TestMinus(t *testing.T) {
	a := 4
	b := 1
	expect := 3

	if Minus(a, b) != expect {
		t.Errorf("Expected answer is %v", expect)
	}
}

func TestMultiple(t *testing.T) {
	a := 4
	b := 2
	expect := 8

	if Multiple(a, b) != expect {
		t.Errorf("Expected answer is %v", expect)
	}
}

func TestDivide(t *testing.T) {
	a := 4
	b := 2
	expect := 2

	if Divide(a, b) != expect {
		t.Errorf("Expected answer is %v", expect)
	}
}
