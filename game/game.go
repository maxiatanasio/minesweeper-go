package game

import (
	"github.com/JakeHL/Goid"
	"minesweeper-API/helpers"
)

const (
	_ = IOTA
	EMPTY
	MINE
)

type board = [][]int

type game struct {
	uuid *goid.UUID
	board
}

type Options struct {
	SizeX int
	SizeY int
}

var games []game

func Start(options Options) *goid.UUID {
	newGame := game{
		uuid: goid.NewV4UUID(),
		board: [][]int{
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		},
	}

	return newGame.uuid
}

func createBoard(x int, y int) board {
	newBoard := board{}
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			newBoard[i][j] = helpers.RandomInt(1, 2)
		}
	}
	return newBoard
}
