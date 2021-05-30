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

func ReadPattern(filePath string) map[Position]int {
	result := make(map[Position]int)
	file, err := os.Open(filePath)

	checkError(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		currentLineLetters := []rune(scanner.Text())
		for col, letter := range currentLineLetters {
			if letter == 'X' {
				result[Position{ X: row, Y: col }] = 0
			}
			col += 1
		}
		row += 1
	}

	checkError(scanner.Err())

	return result
}
