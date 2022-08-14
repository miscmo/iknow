package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miscmo/iknow/dao"
	"github.com/miscmo/iknow/model"
	"net/http"
)

func AddNodes(c *gin.Context) {
	req := model.AddNodeReq{}
	rsp := model.AddNodeRsp{}

	body := c.Request.Body

	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&req); err != nil {
		rsp.Code = model.ParamInvild
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		fmt.Printf("decode failed, err: %+v, rsp: %v\n", err, rsp)
		return
	}

	fmt.Printf("req: %+v", req)

	if err := dao.InsertNode(req.Node); err != nil {
		rsp.Code = model.UpdateDBFailed
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		fmt.Printf("decode failed, rsp: %v\n", rsp)
		return
	}

	c.JSON(http.StatusOK, rsp)
}
