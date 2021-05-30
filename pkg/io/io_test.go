package io

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGivenAPatternFileReadPatternReturnsHashOfAliveCells( t *testing.T ) {
	expectedResultHashOfAliveCells := map[Position]int {
		Position { X:1, Y:1 }: 0,
	}
	expectedResultMinBorder := Position{X: 0, Y: 0}
	expectedResultMaxBorder := Position{X: 3, Y: 3}
	pathFile := "../../patterns/test_pattern.txt"
	resultHashOfAliveCells, resultMinBorder, resultMaxBorder := ReadPattern(pathFile)

	assert.Equal(t, expectedResultHashOfAliveCells, resultHashOfAliveCells)
	assert.Equal(t, expectedResultMinBorder, resultMinBorder)
	assert.Equal(t, expectedResultMaxBorder, resultMaxBorder)
}

func TestGivenHashOfAliveCellsMinBorderAndMaxBorderPrintBoardWriteBoardToConsole( t *testing.T ) {
	expectedResult := "---\n-X-\n---\n"
	hashOfAliveCells := map[Position]int {
		Position{ X:1, Y:1 }: 0,
	}
	minBorder := Position{X: 0, Y: 0}
	maxBorder := Position{X: 3, Y: 3}

	result := PrintBoard(hashOfAliveCells, minBorder, maxBorder)

	assert.Equal(t, expectedResult, result)
}
