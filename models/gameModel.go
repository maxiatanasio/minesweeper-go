package models

import "minesweeper-API/helpers"

type Game struct {
	ID    uint         `gorm:"primary_key" json:"id"`
	Uuid  string       `json:"uuid"`
	Board helpers.JSON `sql:"type:json" json:"board"`
}
