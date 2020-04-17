package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maxiatanasio/mineswepper-API/controllers"
	"net/http"
)

func main() {

	db := ConfigDB()

	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
		})
	})

	r.GET("/game/start/:x/:y", controllers.CreateGame(db))
	r.GET("/game/status/:uuid", controllers.GameStatus(db))
	r.GET("/game/click/:uuid/:x/:y", controllers.GameClick(db))
	r.GET("/game/draw/:uuid", controllers.GameDraw(db))
	r.GET("/game/flag/:uuid/:x/:y", controllers.GameFlag(db))

	r.Run("localhost:4657")

	defer db.Close()
}
