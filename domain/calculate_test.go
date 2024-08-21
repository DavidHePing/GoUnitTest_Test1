package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	a := 1
	b := 2
	expect := 3

	assert.Equal(t, Add(a, b), expect)
}

func TestMinus(t *testing.T) {
	a := 4
	b := 1
	expect := 3

	assert.Equal(t, Minus(a, b), expect)
}

func TestMultiple(t *testing.T) {
	a := 4
	b := 2
	expect := 8

	assert.Equal(t, Multiple(a, b), expect)
}

func TestDivide(t *testing.T) {
	a := 4
	b := 2
	expect := 2

	assert.Equal(t, Divide(a, b), expect)
}
