package gameService

import (
	"encoding/json"
	goid "github.com/JakeHL/Goid"
	"github.com/jinzhu/gorm"
	"minesweeper-API/models"
)

func SearchGame(uuid string, db *gorm.DB) (*Game, error) {

	var game models.Game

	if result := db.Where(&models.Game{
		Uuid: uuid,
	}).First(&game); result.Error != nil {
		return nil, result.Error
	}

	uuidFromModel, _ := goid.GetUUIDFromString(game.Uuid)

	var board Board
	var mines MineStats

	json.Unmarshal(game.Board, &board)
	json.Unmarshal(game.Mines, &mines)

	gameFounded := Game{
		uuid:   uuidFromModel,
		Status: game.Status,
		Board:  board,
		Mines:  mines,
	}

	return &gameFounded, nil

}
