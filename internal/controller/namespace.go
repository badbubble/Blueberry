package controller

import (
	"Blueberry/internal/logic"
	"github.com/gin-gonic/gin"
)

// GetNamespace godoc
// @Summary Get all available namespaces of Kubernetes
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /hello [get]
func GetNamespace(c *gin.Context) {
	err, namespaceList := logic.GetNamespaceList()
	if err != nil {
		return
	}
	ResponseSuccess(c, namespaceList)
}
