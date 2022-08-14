package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miscmo/iknow/dao"
	"github.com/miscmo/iknow/proto"
	"net/http"
)

func AddNode(c *gin.Context) {
	req := proto.AddNodeReq{}
	rsp := proto.AddNodeRsp{}

	body := c.Request.Body

	decoder := json.NewDecoder(body)

	//var req SearchNodeReq
	if err := decoder.Decode(&req); err != nil {
		rsp.Code = proto.RspCode_ParamInvild
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		fmt.Printf("decode failed, rsp: %v\n", rsp)
		return
	}

	if err := dao.InsertManyNode(req.Nodes); err != nil {
		rsp.Code = proto.RspCode_UpdateDBFailed
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		fmt.Printf("decode failed, rsp: %v\n", rsp)
		return
	}

	c.JSON(http.StatusOK, rsp)
}
