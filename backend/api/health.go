package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "health ok"})

}
