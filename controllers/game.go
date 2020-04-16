package controllers

import (
	"github.com/gin-gonic/gin"
	"minesweeper-API/gameService"
	"net/http"
	"strconv"
)

func CreateGame(c *gin.Context) {

	x, err := strconv.Atoi(c.Param("x"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameters sent",
		})
		return
	}

	y, err := strconv.Atoi(c.Param("y"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameters sent",
		})
		return
	}

	uuid := gameService.Start(gameService.Options{
		SizeX: x,
		SizeY: y,
	})

	c.JSON(http.StatusOK, gin.H{
		"uuid": uuid.String(),
	})
}

func GameStatus(c *gin.Context) {
	uuid := c.Param("uuid")

	gameStatus, err := gameService.Status(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"game": gameStatus,
	})

}

func GameClick(c *gin.Context) {
	x, err := strconv.Atoi(c.Param("x"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameters sent",
		})
		return
	}

	y, err := strconv.Atoi(c.Param("y"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameters sent",
		})
		return
	}

	uuid := c.Param("uuid")

	game, err := gameService.Click(uuid, x, y)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"game": game,
	})

}

func GameDraw(c *gin.Context) {
	uuid := c.Param("uuid")

	response, err := gameService.Draw(uuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, *response)

}

func GameFlag(c *gin.Context) {
	x, err := strconv.Atoi(c.Param("x"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameters sent",
		})
		return
	}

	y, err := strconv.Atoi(c.Param("y"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Parameters sent",
		})
		return
	}

	uuid := c.Param("uuid")

	game, err := gameService.Flag(uuid, x, y)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"game": game,
	})

}
