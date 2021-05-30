package game

import (
	"github.com/ayaar25/golang-game-of-life/pkg/io"
	"os/exec"
	"os"
	"time"
	"fmt"
)

func Entry(filePath string) {
	hashOfAliveCells, minBorder, maxBorder := io.ReadPattern(filePath)
	board := Board{hashOfAliveCells, minBorder, maxBorder}

	generation := 0
	for true {
		time.Sleep(200000000 * time.Nanosecond)
	
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println("Generation: ", generation)
		io.PrintBoard(board.hashOfAliveCells, board.minBorder, board.maxBorder)
		board.GenerateNextGenerationOfAliveCells()
		board.RenewBorders()
		generation += 1
	}
}
