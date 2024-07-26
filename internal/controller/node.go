package controller

import (
	"Blueberry/internal/logic"
	"Blueberry/internal/model"
	"github.com/gin-gonic/gin"
)

func GetNodeHandler(c *gin.Context) {
	nodeName := c.Query("nodeName")
	// get list
	if nodeName == "" {
		nodeList, err := logic.GetNodeList()
		if err != nil {
			ResponseError(c, CodeGetNodeListError)
			return
		}
		ResponseSuccess(c, nodeList)
	}
	// get the detail of a node
	nodeDetail, err := logic.GetNodeDetail(nodeName)
	if err != nil {
		ResponseError(c, CodeGetNodeListError)
		return
	}
	ResponseSuccess(c, nodeDetail)
}

func UpdateNodeLabelsHandler(c *gin.Context) {
	labelsUpdate := &model.NodeLabelsUpdate{}
	if err := c.ShouldBind(labelsUpdate); err != nil {
		ResponseErrorWithMsg(c, CodeParameterError, err.Error())
		return
	}

	err := logic.UpdateNodeLabels(labelsUpdate)
	if err != nil {
		ResponseErrorWithMsg(c, CodeUpdateNodeLabelsError, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}

func UpdateNodeTaintsHandler(c *gin.Context) {
	taintsUpdate := &model.NodeTaintUpdate{}
	if err := c.ShouldBind(taintsUpdate); err != nil {
		ResponseErrorWithMsg(c, CodeParameterError, err.Error())
		return
	}

	err := logic.UpdateNodeTaints(taintsUpdate)
	if err != nil {
		ResponseErrorWithMsg(c, CodeUpdateNodeLabelsError, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}
