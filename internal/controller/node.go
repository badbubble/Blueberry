package controller

import (
	"Blueberry/internal/logic"
	"github.com/gin-gonic/gin"
)

func GetNodeHandler(c *gin.Context) {
	nodeList, err := logic.GetNodeList()
	if err != nil {
		ResponseError(c, CodeGetNodeListError)
		return
	}
	ResponseSuccess(c, nodeList)

}
