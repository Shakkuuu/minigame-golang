package main

import (
	"fmt"

	"github.com/Shakkuuu/minigame-golang/minigame"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var coin minigame.Coin

func main() {
	fmt.Println("start")

	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")

	r.GET("/", index)
	r.GET("minigame", minigamemain)

	r.POST("addcoin", addcoin)
	r.POST("createspeedup", CreateSpeedUp)

	// 8082ポートで起動
	r.Run(":8084")
}

func index(c *gin.Context) {
	if coin.Speed < 1 {
		coin.Speed = 1
	}
	if coin.Speedneed < 20 {
		coin.Speedneed = 20
	}
	coin = minigame.Coin{Qty: coin.Qty, Speed: coin.Speed, Speedneed: coin.Speedneed}
	c.HTML(200, "index.html", nil)
}

func minigamemain(c *gin.Context) {
	coi := coin
	fmt.Println(coi)
	c.HTML(200, "game.html", gin.H{"Qty": int(coi.Qty), "Speed": coi.Speed, "Speedneed": int(coi.Speedneed)})
}

func addcoin(c *gin.Context) {
	coin.Qty += coin.Speed
	c.Redirect(303, "/minigame")
}

func CreateSpeedUp(c *gin.Context) {
	if int(coin.Qty) < int(coin.Speedneed) {
		c.Redirect(303, "/minigame")
		return
	}

	coin.Qty -= float32(int(coin.Speedneed))
	coin.Speed += 0.4

	coin.Speedneed *= 1.7
	c.Redirect(303, "/minigame")
}
