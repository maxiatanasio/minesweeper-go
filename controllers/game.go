package controllers

import (
	"github.com/gin-gonic/gin"
	"minesweeper-API/game"
	"net/http"
	"strconv"
)

func CreateGame(c *gin.Context) {

	x, err := strconv.Atoi(c.Param("x"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameters sent",
		})
	}

	y, err := strconv.Atoi(c.Param("y"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameters sent",
		})
	}

	uuid := game.Start(game.Options{
		SizeX: x,
		SizeY: y,
	})

	c.JSON(http.StatusOK, gin.H{
		"uuid": uuid.String(),
	})
}
