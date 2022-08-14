package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Calc(c *gin.Context)  {
	num1 := c.Request.FormValue("num1")
	num2 := c.Request.FormValue("num2")
	reqBody := c.Request.Body
	reqjson, _ := io.ReadAll(reqBody)

	fmt.Printf("num1 = %v, num2 = %v, reqJson: %v\n", num1, num2, reqjson)

	c.JSON(http.StatusOK, gin.H{"add": "testadd", "mul": "testmul", "sub": "testsub", "div": "testdiv"})

}
