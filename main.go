package main

import (
	"github.com/gin-gonic/gin"
	"minesweeper-API/controllers"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	r.GET("/game/start/:x/:y", controllers.CreateGame)
	r.GET("/game/status/:uuid", controllers.GameStatus)

	r.Run("localhost:4657")
}
