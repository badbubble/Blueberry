package router

import (
	"Blueberry/internal/controller"
	"Blueberry/internal/middleware"
	"Blueberry/pkg/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.CORSMiddleware(), logger.GinLogger())
	v1 := r.Group("/api/v1")
	{
		k8sGroup := v1.Group("/k8s")
		{
			k8sGroup.GET("/pod", controller.GetPodHandler)
			k8sGroup.POST("/pod", controller.CreatePodHandler)
			k8sGroup.DELETE("/pod", controller.DeletePodHandler)
			k8sGroup.PUT("/pod", controller.UpdatePodHandler)

			k8sGroup.GET("/namespace", controller.GetNamespace)
		}
	}
	return r
}
