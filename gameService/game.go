package gameService

import (
	"github.com/jakehl/goid"
	"github.com/jinzhu/gorm"
	"github.com/maxiatanasio/mineswepper-API/helpers"
	"strconv"
)

func Start(options Options, db *gorm.DB) (*goid.UUID, error) {
	board, mines := createBoard(options.SizeX, options.SizeY)
	newGame := Game{
		uuid:   goid.NewV4UUID(),
		Board:  board,
		Status: inProgress,
		Mines:  mines,
	}

	if err := CreateGame(&newGame, db); err != nil {
		return nil, err
	}

	return newGame.uuid, nil
}

func Status(uuid string, db *gorm.DB) (*Game, error) {
	game, _, err := SearchGame(uuid, db)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func Click(uuid string, x int, y int, db *gorm.DB) (*Game, error) {
	game, id, err := SearchGame(uuid, db)
	if err != nil {
		return nil, err
	}

	if game.Status != inProgress {
		return game, nil
	}

	if game.Board[x][y].Status == nothing {
		game.Board[x][y].Status = open
		game.Mines.Discovered++
		if game.Board[x][y].Mine == mine {
			game.Status = lost
			if err := UpdateGame(game, id, db); err != nil {
				return nil, err
			}
			return game, nil
		}

		if game.Board[x][y].Adyacents == 0 {
			clickCellEvent(game, x, y)
		}

		if game.Mines.Discovered == getGameCellCount(game)-game.Mines.Total {
			game.Status = won
		}

		if err := UpdateGame(game, id, db); err != nil {
			return nil, err
		}

	}

	return game, nil

}

func Draw(uuid string, db *gorm.DB) (*string, error) {
	game, _, err := SearchGame(uuid, db)
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
					response = response + strconv.Itoa(cell.Adyacents)
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

func Flag(uuid string, x int, y int, db *gorm.DB) (*Game, error) {
	game, id, err := SearchGame(uuid, db)
	if err != nil {
		return nil, err
	}

	if game.Status != inProgress {
		return game, nil
	}

	if game.Board[x][y].Status == nothing {
		game.Board[x][y].Status = flagged
		game.Mines.Flagged++
		if err := UpdateGame(game, id, db); err != nil {
			return nil, err
		}
	}

	return game, nil
}

func clickCellEvent(game *Game, x int, y int) (*Game, error) {

	adyacentCells := getAdyacentCells(&game.Board, x, y)

	for i := 0; i < len(adyacentCells); i++ {
		clickWaveEffect(game, adyacentCells[i])
	}

	return game, nil

}

func clickWaveEffect(game *Game, cell *Cell) {
	if cell.Mine == empty && cell.Status == nothing {
		cell.Status = open
		game.Mines.Discovered++
		if cell.Adyacents == 0 {
			clickCellEvent(game, cell.X, cell.Y)
		}
	}
}

func getAdyacentCells(board *Board, x int, y int) []*Cell {
	adyacentCells := []*Cell{}

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

func createBoard(x int, y int) (Board, MineStats) {
	newBoard := Board{}
	mines := MineStats{
		Total:      0,
		Flagged:    0,
		Discovered: 0,
	}
	for i := 0; i < y; i++ {
		newBoard = append(newBoard, []Cell{})
		for j := 0; j < x; j++ {

			mineValue := empty
			randomMineValueBase := helpers.RandomInt(1, 10)
			if randomMineValueBase > 9 {
				mineValue = mine
				mines.Total++
			}

			newBoard[i] = append(newBoard[i], Cell{
				Mine:   mineValue,
				Status: nothing,
				X:      i,
				Y:      j,
			})
		}
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			newBoard[i][j].Adyacents = evaluateAdyacents(&newBoard, i, j)
		}
	}
	return newBoard, mines
}

func evaluateAdyacents(board *Board, x int, y int) int {
	adyacentCells := getAdyacentCells(board, x, y)

	adyacentMinesCounter := 0
	for i := 0; i < len(adyacentCells); i++ {
		if adyacentCells[i].Mine == mine {
			adyacentMinesCounter++
		}
	}

	return adyacentMinesCounter

}

func getGameCellCount(game *Game) int {
	return len(game.Board) * len(game.Board[0])
}
