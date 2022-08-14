package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/miscmo/iknow/dao"
	"github.com/miscmo/iknow/model"
	"net/http"
)

func UpdateNode(c *gin.Context)  {
	req := model.UpdateNodeReq{}
	rsp := model.UpdateNodeRsp{}

	body := c.Request.Body

	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&req); err != nil {
		rsp.Code = model.ParamInvild
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		fmt.Println("decode failed, err: %s, rsp: %+v", err.Error(), rsp)
		return
	}

	var options []dao.ModOption

	if req.Node.Name.Defined {
		options = append(options, dao.SetNodeName(*req.Node.Name.Value))
	}
	if req.Node.Desc.Defined {
		options = append(options, dao.SetNodeDesc(*req.Node.Desc.Value))
	}
	if req.Node.Type.Defined {
		options = append(options, dao.SetNodeType(*req.Node.Type.Value))
	}
	if req.Node.Collapse.Defined {
		options = append(options, dao.SetNodeCollapse(*req.Node.Collapse.Value))
	}
	if req.Node.Score.Defined {
		options = append(options, dao.SetNodeScore(*req.Node.Score.Value))
	}
	if req.Node.Permission.Defined {
		options = append(options, dao.SetNodePermission(*req.Node.Permission.Value))
	}

	if err := dao.UpdateNode(req.Id, options...); err != nil {
		rsp.Code = model.UpdateDBFailed
		rsp.Msg = err.Error()
		c.JSON(http.StatusOK, rsp)
		return
	}


	fmt.Printf("req: %+v", req)

}
