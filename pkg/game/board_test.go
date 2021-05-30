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

func TestGivenHashOfCellsCountAliveNeighboursEachCellReturnHashOfCellsWithNumberOfAliveNeighbours( t *testing.T ) {
	hashOfAliveCells := map[io.Position]int {
		io.Position{X: 1, Y: 1}: 0,
		io.Position{X: 2, Y: 1}: 0,
	}
	minBorder := io.Position{X: 0, Y: 0}
	maxBorder := io.Position{X: 3, Y: 3}
	expectedResult := map[io.Position]int {
		io.Position{X: 1, Y: 1}: 1,
		io.Position{X: 2, Y: 1}: 1,
	}

	board := Board{hashOfAliveCells, minBorder, maxBorder}
	
	result := board.CountAliveNeighboursEachCell(hashOfAliveCells)
	
	assert.Equal(t, expectedResult, result)	
}

func TestGivenACellAndItsNumberOfAliveNeighboursIsOneIsCellAliveInNextGenerationReturnFalse( t *testing.T ) {
	hashOfAliveCells := map[io.Position]int {
		io.Position{X: 0, Y: 1}: 1,
		io.Position{X: 1, Y: 1}: 1,
	}
	minBorder := io.Position{X: 0, Y: 0}
	maxBorder := io.Position{X: 3, Y: 3}

	board := Board{hashOfAliveCells, minBorder, maxBorder}

	cell := io.Position{X: 0, Y: 1}
	numberOfAliveNeighbours := board.hashOfAliveCells[cell]

	result := board.IsCellAliveInNextGeneration(cell, numberOfAliveNeighbours)
	
	assert.False(t, result)
}

func TestGivenACellAndItsNumberOfAliveNeighboursIsTwoIsCellAliveInNextGenerationReturnTrue( t *testing.T ) {
	hashOfAliveCells := map[io.Position]int {
		io.Position{X: 0, Y: 1}: 2,
		io.Position{X: 1, Y: 1}: 1,
		io.Position{X: 0, Y: 2}: 1,
	}
	minBorder := io.Position{X: 0, Y: 0}
	maxBorder := io.Position{X: 3, Y: 3}

	board := Board{hashOfAliveCells, minBorder, maxBorder}

	cell := io.Position{X: 0, Y: 1}
	numberOfAliveNeighbours := board.hashOfAliveCells[cell]

	result := board.IsCellAliveInNextGeneration(cell, numberOfAliveNeighbours)
	
	assert.True(t, result)
}
