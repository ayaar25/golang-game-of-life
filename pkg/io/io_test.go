package io

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGivenAPatternFileReadPatternReturnsHashOfAliveCells( t *testing.T ) {
	expectedResult := map[Position]int {
		Position { X:1, Y:1 }: 0,
	}
	pathFile := "../../patterns/test_pattern.txt"
	result := ReadPattern(pathFile)

	assert.Equal(t, expectedResult, result)
}
