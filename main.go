package main

import (
	"github.com/gin-gonic/gin"
)

func mainPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"Blog":   "www.miscmo.org",
		"wechat": "miscmo_org",
	})
}

func musicList(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func main() {
	r := gin.Default()
	r.GET("/", mainPage)
	r.GET("/music", musicList)
	r.Run(":8080")
}
