package gameService

import (
	"errors"
	"github.com/JakeHL/Goid"
	"minesweeper-API/helpers"
	"strconv"
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

const (
	_ = iota
	inProgress
	won
	lost
)

type mineStats struct {
	Total   int `json:"total"`
	Flagged int `json:"flagged"`
}

type cell struct {
	Mine      int `json:"mine"`
	Status    int `json:"status"`
	adyacents int
	x         int
	y         int
}

type board = [][]cell

type game struct {
	uuid   *goid.UUID
	Board  board     `json:"board"`
	Status int       `json:"status"`
	Mines  mineStats `json:"mines"`
}

type Options struct {
	SizeX int
	SizeY int
}

var games []game

func Start(options Options) *goid.UUID {
	board, mines := createBoard(options.SizeX, options.SizeY)
	newGame := game{
		uuid:   goid.NewV4UUID(),
		Board:  board,
		Status: inProgress,
		Mines:  mines,
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

func Click(uuid string, x int, y int) (*game, error) {
	game, err := searchGame(uuid)
	if err != nil {
		return nil, err
	}

	game.Board[x][y].Status = open
	if game.Board[x][y].Mine == mine {
		game.Status = lost
		return game, nil
	}

	if game.Board[x][y].adyacents == 0 {
		return clickCellEvent(game, x, y)
	}

	return game, nil

}

func Draw(uuid string) (*string, error) {
	game, err := searchGame(uuid)
	if err != nil {
		return nil, err
	}

	response := ""

	for i := 0; i < len(game.Board); i++ {
		for j := 0; j < len(game.Board[0]); j++ {
			cell := game.Board[i][j]
			if cell.Status == nothing {
				response = response + "?"
			} else if cell.Status == flagged {
				response = response + "F"
			} else if cell.Status == open {
				if cell.Mine == empty {
					response = response + strconv.Itoa(cell.adyacents)
				} else {
					response = response + "M"
				}
			}

		}
		response = response + "\n\r"
	}

	response = response + "\n\rStatus:" + strconv.Itoa(game.Status)
	response = response + "\n\rTotal Mines:" + strconv.Itoa(game.Mines.Total)
	response = response + "\n\rFlagged Mines:" + strconv.Itoa(game.Mines.Flagged)

	return &response, nil
}

func clickCellEvent(game *game, x int, y int) (*game, error) {

	adyacentCells := getAdyacentCells(&game.Board, x, y)

	for i := 0; i < len(adyacentCells); i++ {
		clickWaveEffect(game, adyacentCells[i])
	}

	return game, nil

}

func clickWaveEffect(game *game, cell *cell) {
	if cell.Mine == empty {
		cell.Status = open
		if cell.adyacents == 0 {
			clickCellEvent(game, cell.x, cell.y)
		}
	}
}

func getAdyacentCells(board *board, x int, y int) []*cell {
	adyacentCells := []*cell{}

	if y > 0 {
		if (*board)[x][y-1].Status == nothing {
			adyacentCells = append(adyacentCells, &((*board)[x][y-1]))
		}
	}

	if x > 0 {
		if (*board)[x-1][y].Status == nothing {
			adyacentCells = append(adyacentCells, &((*board)[x-1][y]))
		}
	}

	if y < len((*board)[0])-1 {
		if (*board)[x][y+1].Status == nothing {
			adyacentCells = append(adyacentCells, &((*board)[x][y+1]))
		}
	}

	if x < len(*board)-1 {
		if (*board)[x+1][y].Status == nothing {
			adyacentCells = append(adyacentCells, &((*board)[x+1][y]))
		}
	}

	if y > 0 && x > 0 {
		if (*board)[x-1][y-1].Status == nothing {
			adyacentCells = append(adyacentCells, &((*board)[x-1][y-1]))
		}
	}

	if y > 0 && x < len(*board)-1 {
		if (*board)[x+1][y-1].Status == nothing {
			adyacentCells = append(adyacentCells, &((*board)[x+1][y-1]))
		}
	}

	if y < len((*board)[0])-1 && x < len(*board)-1 {
		if (*board)[x+1][y+1].Status == nothing {
			adyacentCells = append(adyacentCells, &((*board)[x+1][y+1]))
		}
	}

	if y < len((*board)[0])-1 && x > 0 {
		if (*board)[x-1][y+1].Status == nothing {
			adyacentCells = append(adyacentCells, &((*board)[x-1][y+1]))
		}
	}

	return adyacentCells
}

func searchGame(uuid string) (*game, error) {
	for i := 0; i < len(games); i++ {
		if games[i].uuid.String() == uuid {
			return &games[i], nil
		}
	}
	return nil, errors.New("No game found")
}

func createBoard(x int, y int) (board, mineStats) {
	newBoard := board{}
	mines := mineStats{
		Total:   0,
		Flagged: 0,
	}
	for i := 0; i < y; i++ {
		newBoard = append(newBoard, []cell{})
		for j := 0; j < x; j++ {

			mineValue := empty
			randomMineValueBase := helpers.RandomInt(1, 6)
			if randomMineValueBase > 5 {
				mineValue = mine
				mines.Total++
			}

			newBoard[i] = append(newBoard[i], cell{
				Mine:   mineValue,
				Status: nothing,
				x:      i,
				y:      j,
			})
		}
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			newBoard[i][j].adyacents = evaluateAdyacents(&newBoard, i, j)
		}
	}
	return newBoard, mines
}

func evaluateAdyacents(board *board, x int, y int) int {
	adyacentCells := getAdyacentCells(board, x, y)

	adyacentMinesCounter := 0
	for i := 0; i < len(adyacentCells); i++ {
		if adyacentCells[i].Mine == mine {
			adyacentMinesCounter++
		}
	}

	return adyacentMinesCounter

}
