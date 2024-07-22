package controller

import (
	"Blueberry/internal/logic"
	"github.com/gin-gonic/gin"
)

func GetNamespace(c *gin.Context) {
	err, namespaceList := logic.GetNamespaceList()
	if err != nil {
		return
	}
	ResponseSuccess(c, namespaceList)
}
