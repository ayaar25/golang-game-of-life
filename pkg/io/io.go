package io

import (
	"os"
	"bufio"
	"log"
)

type Position struct {
	X, Y int
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadPattern(filePath string) (map[Position]int, Position, Position) {
	result := make(map[Position]int)
	file, err := os.Open(filePath)

	checkError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	maxRow := 0
	maxCol := 0
	for scanner.Scan() {
		currentLineLetters := []rune(scanner.Text())
		for col, letter := range currentLineLetters {
			if letter == 'X' {
				result[Position{ X: row, Y: col }] = 0
			}
			col += 1
			maxCol = col
		}
		row += 1
	}

	maxRow = row

	checkError(scanner.Err())

	return result, Position{X: 0, Y: 0}, Position{X: maxRow, Y: maxCol}
}
