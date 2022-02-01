package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedCharacters1(t *testing.T) {

	input, expected := "ABCA", "A"

	result := firstRepeatedCharacter(input)

	assert.Equal(t, *result, expected, fmt.Sprintf("should be %v", expected))
}

func TestRepeatedCharacters2(t *testing.T) {

	input, expected := "BCABA", "B"

	result := firstRepeatedCharacter(input)

	assert.Equal(t, *result, expected, fmt.Sprintf("should be %v", expected))
}

func TestRepeatedCharacters3(t *testing.T) {

	input, expected := "glovol", "l"

	result := firstRepeatedCharacter(input)

	assert.Equal(t, *result, expected, fmt.Sprintf("should be %v", expected))
}

func TestRepeatedCharacters4(t *testing.T) {

	input := "ABC"

	result := firstRepeatedCharacter(input)

	assert.Nil(t, result)
}
