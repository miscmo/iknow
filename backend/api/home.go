package api

import "github.com/gin-gonic/gin"

func HomePage(c *gin.Context) {
	c.HTML(200, "index.html",gin.H{"title":"welcom to iknow note"})
}
