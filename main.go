package main

import (
	"fmt"

	"github.com/Shakkuuu/minigame-golang/minigame"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("start")

	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")

	r.GET("/", index)
	r.GET("minigame", minigamemain)

	r.POST("addcoin", addcoin)

	// 8082ポートで起動
	r.Run(":8084")
}

var coin minigame.Coin

func index(c *gin.Context) {
	coin = minigame.Coin{Qty: coin.Qty, CreateSpeed: 1}
	c.HTML(200, "index.html", nil)
}

func minigamemain(c *gin.Context) {
	coi := coin
	c.HTML(200, "game.html", gin.H{"Coin": coi})
}

func addcoin(c *gin.Context) {
	coin.Qty += coin.CreateSpeed
	c.Redirect(303, "minigame")
}
