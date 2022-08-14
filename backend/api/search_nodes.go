package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miscmo/iknow/dao"
	"github.com/miscmo/iknow/model"
	"net/http"
)

//type SearchNodeReq struct {
//	Id string `json:"id"`
//	Page int `json:"page"`
//	PageSize int `json:"page_size"`
//}
//
//type SearchNodeRsp struct {
//}

func SearchNodes(c *gin.Context)  {
	//id := c.Request.FormValue("id")
	//if id == "" {
	//	c.JSON(http.StatusUnauthorized, gin.H{"err": utils.RspInvaildParam.Msg()})
	//	return
	//}

	req := model.SearchNodesReq{}
	rsp := model.SearchNodesRsp{}

	body := c.Request.Body

	decoder := json.NewDecoder(body)

	//var req SearchNodeReq
	if err := decoder.Decode(&req); err != nil {
		rsp.Code = model.ParamInvild
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		fmt.Printf("decode failed, err: %v\n", rsp)
		return
	}

	//if req.Id == "" {
	//	c.JSON(http.StatusUnauthorized, gin.H{"err": utils.RspInvaildParam.Msg()})
	//	fmt.Printf("req: %v\n", req)
	//	return
	//}

	node, err := dao.SearchNode(req.Id, req.Page, req.PageSize)
	if err != nil {
		rsp.Code = model.UpdateDBFailed
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		fmt.Printf("getnode failed, err: %v\n", rsp)
		return
	}

	rsp.Nodes = node

	fmt.Printf("get node info: %+v",node)

	c.JSON(http.StatusOK, rsp)
}
