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

	// 8082ポートで起動
	r.Run(":8084")
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func minigamemain(c *gin.Context) {
	p := minigame.UserStatus{HP: 10, ATK: 1}
	e := minigame.EnemyStatus{Name: "スライム", HP: 5, ATK: 1}
	c.HTML(200, "game.html", gin.H{"player": p, "suraimu": e})
}
