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

	r.POST("dmgp", dmgplayer)
	r.POST("dmge", dmgenemy)

	// 8082ポートで起動
	r.Run(":8084")
}

func index(c *gin.Context) {
	p = minigame.UserStatus{HP: 10, ATK: 1}
	e = minigame.EnemyStatus{Name: "スライム", HP: 5, ATK: 1}
	c.HTML(200, "index.html", nil)
}

var p minigame.UserStatus
var e minigame.EnemyStatus

func minigamemain(c *gin.Context) {
	player := p
	enemy := e
	c.HTML(200, "game.html", gin.H{"player": player, "suraimu": enemy})
}

func dmgplayer(c *gin.Context) {
	p.HP -= e.ATK
	c.Redirect(303, "/minigame")
}

func dmgenemy(c *gin.Context) {
	e.HP -= p.ATK
	c.Redirect(303, "minigame")
}
