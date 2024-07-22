package router

import (
	"Blueberry/internal/controller"
	"Blueberry/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CORSMiddleware())
	v1 := r.Group("/api/v1")
	{
		k8sGroup := v1.Group("/k8s")
		{
			k8sGroup.GET("/pod", controller.GetPodListHandler)
			k8sGroup.POST("/pod", controller.CreateOrUpdatePod)

			k8sGroup.GET("/namespace", controller.GetNamespace)
		}
	}
	return r
}
