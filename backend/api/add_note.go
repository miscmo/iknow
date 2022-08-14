package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miscmo/iknow/model"
	"net/http"
)

func AddNote(c *gin.Context)  {
	req := model.AddNodeReq{}
	rsp := model.AddNodeRsp{}

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
