package models

import (
	"github.com/maxiatanasio/mineswepper-API/helpers"
	"time"
)

type Game struct {
	ID        uint         `gorm:"primary_key" json:"id"`
	Uuid      string       `json:"uuid"`
	Board     helpers.JSON `sql:"type:text" json:"board"`
	Status    int          `json:"status"`
	Mines     helpers.JSON `sql:"type:text" json:"mines"`
	CreatedAt time.Time    `json:"created_at"`
}
