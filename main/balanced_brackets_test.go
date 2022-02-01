package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBalanced1(t *testing.T) {

	input := "()[]{}"

	result := isBalanced(input)

	assert.Equal(t, true, result)
}

func TestIsBalanced2(t *testing.T) {

	input := "([)])"

	result := isBalanced(input)

	assert.Equal(t, result, false)
}

func TestIsBalanced11(t *testing.T) {

	input := "(((((((((())))))))))"

	result := isBalanced(input)

	assert.Equal(t, result, true)
}
