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

	for aliveCell, numberOfAliveNeighbours := range board.hashOfAliveCells {
		result[aliveCell] = numberOfAliveNeighbours
		for _, neighbourPosition := range NeighboursPositions {
			x := aliveCell.X + neighbourPosition.X
			y := aliveCell.Y + neighbourPosition.Y
			
			result[io.Position{X: x, Y: y}] = 0
		}
	}

	return result
}

func (board *Board) CountAliveNeighboursEachCell(hashOfCells map[io.Position]int) map[io.Position]int {
	result := make(map[io.Position]int)

	for cell, numberOfAliveNeighbours := range hashOfCells {
		totalAliveNeighbours := numberOfAliveNeighbours
		for _, neighbourPosition := range NeighboursPositions {
			x := cell.X + neighbourPosition.X
			y := cell.Y + neighbourPosition.Y

			if _, keyExists := board.hashOfAliveCells[io.Position{X: x, Y: y}]; keyExists {
				totalAliveNeighbours += 1
			}
		}

		result[cell] = totalAliveNeighbours
	}

	return result
}

func (board *Board) IsCellAliveInNextGeneration(cell io.Position, numberOfAliveNeighbours int) bool {
	var result bool

	if _, keyExists := board.hashOfAliveCells[cell]; keyExists {
		result = numberOfAliveNeighbours == 2 || numberOfAliveNeighbours == 3
	}

	return result
}
