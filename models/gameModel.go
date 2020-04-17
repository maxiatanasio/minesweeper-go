package models

import (
	"github.com/maxiatanasio/mineswepper-API/helpers"
	"time"
)

type Game struct {
	ID        uint         `gorm:"primary_key" json:"id"`
	Uuid      string       `json:"uuid"`
	Board     helpers.JSON `sql:"type:json" json:"board"`
	Status    int          `json:"status"`
	Mines     helpers.JSON `sql:"type:json json:"mines"`
	CreatedAt time.Time    `gorm:"default:CURRENT_TIMESTAMP"`
}
