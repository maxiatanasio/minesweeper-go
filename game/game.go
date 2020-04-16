package game

import (
	"errors"
	"github.com/JakeHL/Goid"
	"minesweeper-API/helpers"
)

const (
	_ = iota
	empty
	mine
)

const (
	_ = iota
	nothing
	flagged
	open
)

type cell struct {
	Mine   int `json:"mine"`
	Status int `json:"status"`
}

type board = [][]cell

type game struct {
	uuid  *goid.UUID
	Board board `json:"board"`
}

type Options struct {
	SizeX int
	SizeY int
}

var games []game

func Start(options Options) *goid.UUID {
	newGame := game{
		uuid:  goid.NewV4UUID(),
		Board: createBoard(options.SizeX, options.SizeY),
	}

	games = append(games, newGame)

	return newGame.uuid
}

func Status(uuid string) (*game, error) {
	game, err := searchGame(uuid)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func searchGame(uuid string) (*game, error) {
	for i := 0; i < len(games); i++ {
		if games[i].uuid.String() == uuid {
			return &games[i], nil
		}
	}
	return nil, errors.New("No game found")
}

func createBoard(x int, y int) board {
	newBoard := board{}
	for i := 0; i < y; i++ {
		newBoard = append(newBoard, []cell{})
		for j := 0; j < x; j++ {
			newBoard[i] = append(newBoard[i], cell{
				Mine:   helpers.RandomInt(1, 2),
				Status: nothing,
			})
		}
	}
	return newBoard
}
