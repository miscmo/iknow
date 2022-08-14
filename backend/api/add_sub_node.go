package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miscmo/iknow/model"
	"net/http"
)

func AddSubNode(c *gin.Context)  {
	req := model.AddSubNodeReq{}
	rsp := model.AddSubNodeRsp{}

	body := c.Request.Body

	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&req); err != nil {
		rsp.Code = model.ParamInvild
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		fmt.Println("decode failed, err: %s, rsp: %+v", err.Error(), rsp)
		return
	}

	fmt.Printf("req: %+v", req)

}
