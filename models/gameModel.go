package models

import "minesweeper-API/helpers"

type Game struct {
	ID     uint         `gorm:"primary_key" json:"id"`
	Uuid   string       `json:"uuid"`
	Board  helpers.JSON `sql:"type:json" json:"board"`
	Status int          `json:"status"`
	Mines  helpers.JSON `sql:"type:json json:"mines"`
}
