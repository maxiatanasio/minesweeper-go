package gameService

import goid "github.com/JakeHL/Goid"

type MineStats struct {
	Total      int `json:"total"`
	Flagged    int `json:"flagged"`
	Discovered int `json:"discovered"`
}

type Cell struct {
	Mine      int `json:"mine"`
	Status    int `json:"status"`
	Adyacents int `json:"adyacents"`
	X         int `json:"x"`
	Y         int `json:"y"`
}

type Board = [][]Cell

type Game struct {
	uuid   *goid.UUID
	Board  `json:"board"`
	Status int       `json:"status"`
	Mines  MineStats `json:"mines"`
}

type Options struct {
	SizeX int
	SizeY int
}
