package main

import (
	"github.com/gin-gonic/gin"
	"github.com/miscmo/soil/blog"
	"github.com/miscmo/soil/music"
)

func MainPage(c *gin.Context) {
	c.HTML(200, "index.html", "")
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("ui/*.html")
	r.Static("/static", "./ui/static")
	r.GET("/", MainPage)
	r.GET("/MusicHome", music.MusicHome)
	r.GET("/MusicList", music.MusicList)
	r.POST("/AddMusic", music.AddMusic)
	r.GET("/PlayMusic/:id", music.PlayMusic)
	r.GET("/BlogHome", blog.BlogHome)
	r.GET("/BlogList", blog.BlogList)
	r.Run(":8080")
}
