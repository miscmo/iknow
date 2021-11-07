package blog

import (
	"time"

	"github.com/gin-gonic/gin"
)

type BlogInfo struct {
	Id         int64
	Title      string
	Author     string
	Desc       string
	CreateTime time.Time
	ModTime    time.Time
}

func BlogHome(c *gin.Context) {
	c.HTML(200, "blog_index.html", "")
}

func BlogList(c *gin.Context) {
	c.JSON(200, []BlogInfo{})
}
