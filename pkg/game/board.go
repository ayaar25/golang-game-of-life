package game

import (
	"github.com/ayaar25/golang-game-of-life/pkg/io"
)

type Board struct {
	hashOfAliveCells map[io.Position]int
	minBorder io.Position
	maxBorder io.Position
}

var NeighboursPositions = map[string]io.Position {
	"LEFT": io.Position{X: 0, Y: -1},
	"TOP_LEFT": io.Position{X: -1, Y: -1},
	"TOP": io.Position{X: -1, Y: 0},
	"TOP_RIGHT": io.Position{X: -1, Y: 1},
	"RIGHT": io.Position{X: 0, Y: 1},
	"BOTTOM_RIGHT": io.Position{X: 1, Y: 1},
	"BOTTOM": io.Position{X: 1, Y: 0},
	"BOTTOM_LEFT": io.Position{X: 1, Y: -1},
}

func (board *Board) GenerateHashOfAliveCellsAndItsNeighbours() map[io.Position]int {
	result := make(map[io.Position]int)

	for aliveCell, numberOfNeighbours := range board.hashOfAliveCells {
		result[aliveCell] = numberOfNeighbours
		for _, neighbourPosition := range NeighboursPositions {
			x := aliveCell.X + neighbourPosition.X
			y := aliveCell.Y + neighbourPosition.Y
			
			result[io.Position{X: x, Y: y}] = 0
		}
	}

	return result
}
