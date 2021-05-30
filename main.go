package main

import (
	"github.com/ayaar25/golang-game-of-life/pkg/game"
	"os"
)

func main() {
	filePath := os.Args[1]

	game.Entry(filePath)
}
