package io

import (
	"os"
	"bufio"
	"log"
	"strings"
	"fmt"
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

func PrintBoard(hashOfAliveCells map[Position]int, minBorder Position, maxBorder Position) string {
	var printResult strings.Builder

	for row := minBorder.X; row < maxBorder.X; row++ {
		for col := minBorder.Y; col < maxBorder.Y; col++ {
			if _, keyExists := hashOfAliveCells[Position{X: row, Y: col}]; keyExists {
				printResult.WriteString("X")
			} else {
				printResult.WriteString("-")
			}
			
			if col == maxBorder.Y - 1 {
				printResult.WriteString("\n")
			}
		}
	}

	fmt.Println(printResult.String())
	return printResult.String()
}
