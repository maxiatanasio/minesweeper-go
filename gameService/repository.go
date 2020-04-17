package gameService

import (
	"github.com/jinzhu/gorm"
	"github.com/maxiatanasio/mineswepper-API/models"
)

func CreateGame(game *Game, db *gorm.DB) error {

	if result := db.Create(dbModelFromGame(game, nil)); result.Error != nil {
		return result.Error
	}

	return nil
}

func SearchGame(uuid string, db *gorm.DB) (*Game, *uint, error) {

	var gameModel models.Game

	if result := db.Where(&models.Game{
		Uuid: uuid,
	}).First(&gameModel); result.Error != nil {
		return nil, nil, result.Error
	}

	game, id := gameFromDBModel(&gameModel)

	return game, id, nil

}

func UpdateGame(game *Game, id *uint, db *gorm.DB) error {
	gameModel := dbModelFromGame(game, id)
	if result := db.Save(gameModel); result.Error != nil {
		return result.Error
	}
	return nil
}
