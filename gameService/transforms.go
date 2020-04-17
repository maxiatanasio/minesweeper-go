package gameService

import (
	"encoding/json"
	goid "github.com/JakeHL/Goid"
	"minesweeper-API/models"
)

func gameFromDBModel(gameModel *models.Game) (*Game, *uint) {
	uuidFromModel, _ := goid.GetUUIDFromString(gameModel.Uuid)

	var board Board
	var mines MineStats

	json.Unmarshal(gameModel.Board, &board)
	json.Unmarshal(gameModel.Mines, &mines)

	gameFounded := Game{
		uuid:   uuidFromModel,
		Status: gameModel.Status,
		Board:  board,
		Mines:  mines,
	}

	return &gameFounded, &gameModel.ID
}

func dbModelFromGame(game *Game, id *uint) *models.Game {
	jsonBoard, _ := json.Marshal(game.Board)
	jsonMines, _ := json.Marshal(game.Mines)

	gameModel := models.Game{
		Uuid:   game.uuid.String(),
		Board:  jsonBoard,
		Status: game.Status,
		Mines:  jsonMines,
	}

	if id != nil {
		gameModel.ID = *id
	}

	return &gameModel
}
