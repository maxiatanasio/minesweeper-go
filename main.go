package main

import (
	"github.com/gin-gonic/gin"
	"minesweeper-API/controllers"
	"net/http"
)

func main() {

	ConfigDB()

	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	r.GET("/game/start/:x/:y", controllers.CreateGame)
	r.GET("/game/status/:uuid", controllers.GameStatus)
	r.GET("/game/click/:uuid/:x/:y", controllers.GameClick)
	r.GET("/game/draw/:uuid", controllers.GameDraw)
	r.GET("/game/flag/:uuid/:x/:y", controllers.GameFlag)

	r.Run("localhost:4657")
}
