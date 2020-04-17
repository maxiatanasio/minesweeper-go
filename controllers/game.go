package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/maxiatanasio/mineswepper-API/gameService"
	"net/http"
	"strconv"
)

func CreateGame(db *gorm.DB) func(ctx *gin.Context) {

	return func(c *gin.Context) {
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

		uuid, err := gameService.Start(gameService.Options{
			SizeX: x,
			SizeY: y,
		}, db)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"uuid": uuid.String(),
		})
	}

}

func GameStatus(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")

		gameStatus, err := gameService.Status(uuid, db)
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
}

func GameClick(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
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

		game, err := gameService.Click(uuid, x, y, db)
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
}

func GameDraw(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		uuid := c.Param("uuid")

		response, err := gameService.Draw(uuid, db)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.String(http.StatusOK, *response)
	}
}

func GameFlag(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
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

		game, err := gameService.Flag(uuid, x, y, db)
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
}
