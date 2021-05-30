package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ayaar25/golang-game-of-life/pkg/io"
)

func TestGivenHashOfAliveCellsGenerateHashOfAliveCellsAndItsNeighboursGeneratesNewHashOfAliveCellsAndItsNeighbours( t *testing.T ) {
	hashOfAliveCells := map[io.Position]int {
		io.Position{X: 1, Y: 1}: 0,
	}
	minBorder := io.Position{X: 0, Y: 0}
	maxBorder := io.Position{X: 3, Y: 3}
	expectedResult := map[io.Position]int {
		io.Position{X: 0, Y: 0}: 0,
		io.Position{X: 0, Y: 1}: 0,
		io.Position{X: 0, Y: 2}: 0,
		io.Position{X: 1, Y: 0}: 0,
		io.Position{X: 1, Y: 1}: 0,
		io.Position{X: 1, Y: 2}: 0,
		io.Position{X: 2, Y: 0}: 0,
		io.Position{X: 2, Y: 1}: 0,
		io.Position{X: 2, Y: 2}: 0,
	}

	board := Board{hashOfAliveCells, minBorder, maxBorder}
	
	result := board.GenerateHashOfAliveCellsAndItsNeighbours()
	
	assert.Equal(t, expectedResult, result)
}